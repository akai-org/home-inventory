package api

import (
	"net/http"

	"github.com/akai-org/home-inventory/internal/db"

	"encoding/json"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Router   chi.Router
	Database db.DB
}

func RegisterRoutes(app *App) {
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)

	app.Router.Get("/health", app.HealthCheck)
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
