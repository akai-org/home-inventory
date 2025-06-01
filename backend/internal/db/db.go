package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type DB interface {
	Ping(ctx context.Context) error
	// Add more methods as needed
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
