package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/akai-org/home-inventory/internal/api"
	"github.com/akai-org/home-inventory/internal/db"
	"github.com/huntclauss/dotenv"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting server...")
	fmt.Println("Loading environment variables from .env file...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Environment variables loaded.")

	dbURL := dotenv.MustGet("DATABASE_URL")
	address := dotenv.GetDefault("ADDRESS", "127.0.0.1:8080")

	fmt.Printf("Connecting to database at %s...\n", dbURL)
	database, err := db.NewPostgres(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	fmt.Println("Connected to database successfully.")

	if err := database.RunMigrations("migrations"); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalf("failed to run migrations: %v", err)
		}
	}
	fmt.Println("Database migrations applied successfully.")

	app := api.App{Router: chi.NewRouter(), Database: database}
	app.Router.Use(middleware.RequestID)
	app.Router.Use(middleware.RealIP)
	app.Router.Use(middleware.Logger)
	app.Router.Use(middleware.Recoverer)
	api.RegisterRoutes(&app)

	fmt.Println("Starting HTTP server...")

	log.Printf("Server starting on %s", address)
	http.ListenAndServe(address, app.Router)
}
