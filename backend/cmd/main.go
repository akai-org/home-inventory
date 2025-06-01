package main

import (
	"log"

	"github.com/akai-org/home-inventory/internal/api"
	"github.com/akai-org/home-inventory/internal/db"

	"net/http"

	"github.com/huntclauss/dotenv"

	"github.com/go-chi/chi/v5"
)

func main() {
	dotenv.LoadEnv(".env")

	dbURL := dotenv.MustGet("DATABASE_URL")
	address := dotenv.GetDefault("ADDRESS", "127.0.0.1:8080")

	database, err := db.NewPostgres(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	app := api.App{Router: chi.NewRouter(), Database: database}
	api.RegisterRoutes(&app)

	http.ListenAndServe(address, app.Router)
}
