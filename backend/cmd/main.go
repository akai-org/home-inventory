package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/akai-org/home-inventory/internal/api"
	"github.com/akai-org/home-inventory/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/huntclauss/dotenv"
)

func main() {
	dotenv.LoadEnv(".env")

	dbURL := dotenv.MustGet("DATABASE_URL")
	address := dotenv.GetDefault("ADDRESS", "127.0.0.1:8080")

	database, err := db.NewPostgres(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := database.RunMigrations("migrations"); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalf("failed to run migrations: %v", err)
		}
	}

	app := api.App{Router: chi.NewRouter(), Database: database}
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
	api.RegisterRoutes(&app)

	log.Printf("Server starting on %s", address)
	http.ListenAndServe(address, app.Router)
}
