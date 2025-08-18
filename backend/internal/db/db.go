package db

import (
	"context"
	"database/sql"
	"io/ioutil"
	"path/filepath"

	"github.com/akai-org/home-inventory/internal/models"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type DB interface {
	Ping(ctx context.Context) error

	// Storage operations
	CreateStorage(ctx context.Context, storage *models.Storage) error
	GetStorage(ctx context.Context, id uuid.UUID) (*models.Storage, error)
	ListStorages(ctx context.Context, parentID *uuid.UUID) ([]*models.Storage, error)
	UpdateStorage(ctx context.Context, storage *models.Storage) error
	DeleteStorage(ctx context.Context, id uuid.UUID) error

	// Item operations
	CreateItem(ctx context.Context, item *models.Item) error
	GetItem(ctx context.Context, id uuid.UUID) (*models.Item, error)
	ListItems(ctx context.Context, storageID *uuid.UUID, filters ItemFilters) ([]*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id uuid.UUID) error

	// Search operations
	SearchItems(ctx context.Context, query string, filters ItemFilters) ([]*models.Item, error)

	// QR Code operations
	CreateQRCode(ctx context.Context, qrCode *models.QRCode) error
	GetQRCode(ctx context.Context, codeData string) (*models.QRCode, error)

	// Migration operations
	RunMigrations(migrationsPath string) error
}

type ItemFilters struct {
	Name      *string
	Type      *string
	Tags      []string
	StorageID *uuid.UUID
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres(url string) (DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (p *Postgres) Ping(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *Postgres) RunMigrations(migrationsPath string) error {
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(migrationsPath, file.Name())
			migration, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			_, err = p.db.Exec(string(migration))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
