package api

import (
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/akai-org/home-inventory/internal/db"
	"github.com/akai-org/home-inventory/internal/models"
	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	qrcode "github.com/skip2/go-qrcode"
)

type App struct {
	Router   chi.Router
	Database db.DB
}

func RegisterRoutes(app *App) {
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
	app.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	app.Router.Get("/health", app.HealthCheck)

	// API routes
	app.Router.Route("/api/v1", func(r chi.Router) {
		// Storage routes
		r.Route("/storages", func(r chi.Router) {
			r.Get("/", app.ListStorages)
			r.Post("/", app.CreateStorage)
			r.Get("/{id}", app.GetStorage)
			r.Put("/{id}", app.UpdateStorage)
			r.Delete("/{id}", app.DeleteStorage)
			r.Get("/{id}/items", app.ListStorageItems)
		})

		// Item routes
		r.Route("/items", func(r chi.Router) {
			r.Get("/", app.ListItems)
			r.Post("/", app.CreateItem)
			r.Get("/{id}", app.GetItem)
			r.Put("/{id}", app.UpdateItem)
			r.Delete("/{id}", app.DeleteItem)
			r.Get("/search", app.SearchItems)
		})

		// QR Code routes
		r.Route("/qr", func(r chi.Router) {
			r.Post("/generate", app.GenerateQRCode)
			r.Get("/{code}", app.ResolveQRCode)
		})

		// Barcode routes
		r.Route("/barcode", func(r chi.Router) {
			r.Post("/generate", app.GenerateBarcode)
		})
	})
}

func (a *App) HealthCheck(w http.ResponseWriter, req *http.Request) {
	if err := a.Database.Ping(req.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "unhealthy"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Storage handlers

func (a *App) ListStorages(w http.ResponseWriter, r *http.Request) {
	var parentID *uuid.UUID
	if parentIDStr := r.URL.Query().Get("parent_id"); parentIDStr != "" {
		id, err := uuid.Parse(parentIDStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid parent_id format", err)
			return
		}
		parentID = &id
	}

	storages, err := a.Database.ListStorages(r.Context(), parentID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list storages", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storages)
}

func (a *App) CreateStorage(w http.ResponseWriter, r *http.Request) {
	var storage models.Storage
	if err := json.NewDecoder(r.Body).Decode(&storage); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	if err := a.Database.CreateStorage(r.Context(), &storage); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create storage", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(storage)
}

func (a *App) GetStorage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid storage ID", err)
		return
	}

	storage, err := a.Database.GetStorage(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "Storage not found", err)
		return	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage)
}

func (a *App) UpdateStorage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid storage ID", err)
		return
	}

	var storage models.Storage
	if err := json.NewDecoder(r.Body).Decode(&storage); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	storage.ID = id
	if err := a.Database.UpdateStorage(r.Context(), &storage); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update storage", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage)
}

func (a *App) DeleteStorage(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid storage ID", err)
		return
	}

	if err := a.Database.DeleteStorage(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete storage", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *App) ListStorageItems(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid storage ID", err)
		return
	}

	filters := parseItemFilters(r)
	items, err := a.Database.ListItems(r.Context(), &id, filters)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list items", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// Item handlers

func (a *App) ListItems(w http.ResponseWriter, r *http.Request) {
	filters := parseItemFilters(r)
	items, err := a.Database.ListItems(r.Context(), nil, filters)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list items", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (a *App) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	if err := a.Database.CreateItem(r.Context(), &item); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create item", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (a *App) GetItem(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid item ID", err)
		return
	}

	item, err := a.Database.GetItem(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "Item not found", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (a *App) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid item ID", err)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	item.ID = id
	if err := a.Database.UpdateItem(r.Context(), &item); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update item", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (a *App) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid item ID", err)
		return
	}

	if err := a.Database.DeleteItem(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete item", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *App) SearchItems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	filters := parseItemFilters(r)

	items, err := a.Database.SearchItems(r.Context(), query, filters)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to search items", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// QR Code handlers

func (a *App) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EntityID   uuid.UUID `json:"entity_id"`
		EntityType string    `json:"entity_type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	if req.EntityType != "item" && req.EntityType != "storage" {
		writeError(w, http.StatusBadRequest, "Entity type must be 'item' or 'storage'", nil)
		return
	}

	// Generate QR code data (URL to entity)
	codeData := req.EntityID.String()

	qrCode := &models.QRCode{
		EntityID:   req.EntityID,
		EntityType: req.EntityType,
		CodeData:   codeData,
	}

	if err := a.Database.CreateQRCode(r.Context(), qrCode); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create QR code", err)
		return
	}

	png, err := qrcode.Encode(codeData, qrcode.Medium, 256)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to generate QR code", err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(png)
}

func (a *App) ResolveQRCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	qrCode, err := a.Database.GetQRCode(r.Context(), code)
	if err != nil {
		writeError(w, http.StatusNotFound, "QR code not found", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(qrCode)
}

// Barcode handlers

func (a *App) GenerateBarcode(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON", err)
		return
	}

	b, err := code128.Encode(req.Data)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to generate barcode", err)
		return
	}

	scaled, err := barcode.Scale(b, 256, 50)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to scale barcode", err)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	png.Encode(w, scaled)
}

// Helper functions

func parseItemFilters(r *http.Request) db.ItemFilters {
	filters := db.ItemFilters{}

	if name := r.URL.Query().Get("name"); name != "" {
		filters.Name = &name
	}

	if itemType := r.URL.Query().Get("type"); itemType != "" {
		filters.Type = &itemType
	}

	if storageIDStr := r.URL.Query().Get("storage_id"); storageIDStr != "" {
		if id, err := uuid.Parse(storageIDStr); err == nil {
			filters.StorageID = &id
		}
	}

	if tags := r.URL.Query()["tags"]; len(tags) > 0 {
		filters.Tags = tags
	}

	return filters
}

func writeError(w http.ResponseWriter, statusCode int, message string, err error) {
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	if err != nil {
		response["details"] = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}
