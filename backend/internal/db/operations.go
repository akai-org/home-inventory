package db

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/akai-org/home-inventory/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Storage operations

func (p *Postgres) CreateStorage(ctx context.Context, storage *models.Storage) error {
	storage.ID = uuid.New()

	query := `
		INSERT INTO storages (id, name, location, description, image_url, parent_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at`

	err := p.db.QueryRowContext(ctx, query,
		storage.ID,
		storage.Name,
		storage.Location,
		storage.Description,
		storage.ImageURL,
		storage.ParentID,
	).Scan(&storage.CreatedAt, &storage.UpdatedAt)

	return err
}

func (p *Postgres) GetStorage(ctx context.Context, id uuid.UUID) (*models.Storage, error) {
	storage := &models.Storage{}

	query := `
		SELECT id, name, location, description, image_url, parent_id, created_at, updated_at
		FROM storages WHERE id = $1`

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&storage.ID,
		&storage.Name,
		&storage.Location,
		&storage.Description,
		&storage.ImageURL,
		&storage.ParentID,
		&storage.CreatedAt,
		&storage.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("storage not found")
	}

	return storage, err
}

func (p *Postgres) ListStorages(ctx context.Context, parentID *uuid.UUID) ([]*models.Storage, error) {
	var query string
	var args []interface{}

	if parentID == nil {
		query = `
			SELECT id, name, location, description, image_url, parent_id, created_at, updated_at
			FROM storages WHERE parent_id IS NULL
			ORDER BY name`
	} else {
		query = `
			SELECT id, name, location, description, image_url, parent_id, created_at, updated_at
			FROM storages WHERE parent_id = $1
			ORDER BY name`
		args = append(args, *parentID)
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var storages []*models.Storage
	for rows.Next() {
		storage := &models.Storage{}
		err := rows.Scan(
			&storage.ID,
			&storage.Name,
			&storage.Location,
			&storage.Description,
			&storage.ImageURL,
			&storage.ParentID,
			&storage.CreatedAt,
			&storage.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		storages = append(storages, storage)
	}

	return storages, rows.Err()
}

func (p *Postgres) UpdateStorage(ctx context.Context, storage *models.Storage) error {
	query := `
		UPDATE storages 
		SET name = $2, location = $3, description = $4, image_url = $5, parent_id = $6
		WHERE id = $1
		RETURNING updated_at`

	err := p.db.QueryRowContext(ctx, query,
		storage.ID,
		storage.Name,
		storage.Location,
		storage.Description,
		storage.ImageURL,
		storage.ParentID,
	).Scan(&storage.UpdatedAt)

	return err
}

func (p *Postgres) DeleteStorage(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM storages WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

// Item operations

func (p *Postgres) CreateItem(ctx context.Context, item *models.Item) error {
	item.ID = uuid.New()

	query := `
		INSERT INTO items (id, name, type, description, image_url, storage_id, tags)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING created_at, updated_at`

	err := p.db.QueryRowContext(ctx, query,
		item.ID,
		item.Name,
		item.Type,
		item.Description,
		item.ImageURL,
		item.StorageID,
		pq.Array(item.Tags),
	).Scan(&item.CreatedAt, &item.UpdatedAt)

	return err
}

func (p *Postgres) GetItem(ctx context.Context, id uuid.UUID) (*models.Item, error) {
	item := &models.Item{}

	query := `
		SELECT id, name, type, description, image_url, storage_id, tags, created_at, updated_at
		FROM items WHERE id = $1`

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Type,
		&item.Description,
		&item.ImageURL,
		&item.StorageID,
		pq.Array(&item.Tags),
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("item not found")
	}

	return item, err
}

func (p *Postgres) ListItems(ctx context.Context, storageID *uuid.UUID, filters ItemFilters) ([]*models.Item, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	query := `
		SELECT id, name, type, description, image_url, storage_id, tags, created_at, updated_at
		FROM items`

	if storageID != nil {
		conditions = append(conditions, fmt.Sprintf("storage_id = $%d", argIndex))
		args = append(args, *storageID)
		argIndex++
	}

	if filters.Name != nil {
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", argIndex))
		args = append(args, "%"+*filters.Name+"%")
		argIndex++
	}

	if filters.Type != nil {
		conditions = append(conditions, fmt.Sprintf("type ILIKE $%d", argIndex))
		args = append(args, "%"+*filters.Type+"%")
		argIndex++
	}

	if len(filters.Tags) > 0 {
		conditions = append(conditions, fmt.Sprintf("tags && $%d", argIndex))
		args = append(args, pq.Array(filters.Tags))
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY name"

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.Description,
			&item.ImageURL,
			&item.StorageID,
			pq.Array(&item.Tags),
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (p *Postgres) UpdateItem(ctx context.Context, item *models.Item) error {
	query := `
		UPDATE items 
		SET name = $2, type = $3, description = $4, image_url = $5, storage_id = $6, tags = $7
		WHERE id = $1
		RETURNING updated_at`

	err := p.db.QueryRowContext(ctx, query,
		item.ID,
		item.Name,
		item.Type,
		item.Description,
		item.ImageURL,
		item.StorageID,
		pq.Array(item.Tags),
	).Scan(&item.UpdatedAt)

	return err
}

func (p *Postgres) DeleteItem(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := p.db.ExecContext(ctx, query, id)
	return err
}

// Search operations

func (p *Postgres) SearchItems(ctx context.Context, query string, filters ItemFilters) ([]*models.Item, error) {
	var conditions []string
	var args []interface{}
	argIndex := 1

	sqlQuery := `
		SELECT id, name, type, description, image_url, storage_id, tags, created_at, updated_at
		FROM items`

	// Add text search condition
	if query != "" {
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR description ILIKE $%d OR type ILIKE $%d)", argIndex, argIndex, argIndex))
		args = append(args, "%"+query+"%")
		argIndex++
	}

	// Apply filters
	if filters.StorageID != nil {
		conditions = append(conditions, fmt.Sprintf("storage_id = $%d", argIndex))
		args = append(args, *filters.StorageID)
		argIndex++
	}

	if filters.Type != nil {
		conditions = append(conditions, fmt.Sprintf("type ILIKE $%d", argIndex))
		args = append(args, "%"+*filters.Type+"%")
		argIndex++
	}

	if len(filters.Tags) > 0 {
		conditions = append(conditions, fmt.Sprintf("tags && $%d", argIndex))
		args = append(args, pq.Array(filters.Tags))
		argIndex++
	}

	if len(conditions) > 0 {
		sqlQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	sqlQuery += " ORDER BY name"

	rows, err := p.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.Description,
			&item.ImageURL,
			&item.StorageID,
			pq.Array(&item.Tags),
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

// QR Code operations

func (p *Postgres) CreateQRCode(ctx context.Context, qrCode *models.QRCode) error {
	qrCode.ID = uuid.New()

	query := `
		INSERT INTO qr_codes (id, entity_id, entity_type, code_data)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at`

	err := p.db.QueryRowContext(ctx, query,
		qrCode.ID,
		qrCode.EntityID,
		qrCode.EntityType,
		qrCode.CodeData,
	).Scan(&qrCode.CreatedAt)

	return err
}

func (p *Postgres) GetQRCode(ctx context.Context, codeData string) (*models.QRCode, error) {
	qrCode := &models.QRCode{}

	query := `
		SELECT id, entity_id, entity_type, code_data, created_at
		FROM qr_codes WHERE code_data = $1`

	err := p.db.QueryRowContext(ctx, query, codeData).Scan(
		&qrCode.ID,
		&qrCode.EntityID,
		&qrCode.EntityType,
		&qrCode.CodeData,
		&qrCode.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("qr code not found")
	}

	return qrCode, err
}
