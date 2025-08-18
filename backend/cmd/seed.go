package main

import (
	"context"
	"log"

	"github.com/akai-org/home-inventory/internal/db"
	"github.com/akai-org/home-inventory/internal/models"
	"github.com/huntclauss/dotenv"
)

func main() {
	dotenv.LoadEnv(".env")

	dbURL := dotenv.MustGet("DATABASE_URL")

	database, err := db.NewPostgres(dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	storage := &models.Storage{
		Name:        "My Room",
		Location:    "On the second floor",
		Description: "The room where I sleep",
	}

	if err := database.CreateStorage(context.Background(), storage); err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}

	items := []models.Item{
		{
			Name:        "Laptop",
			Type:        "Electronic",
			Description: "My personal laptop",
			StorageID:   storage.ID,
			Tags:        []string{"electronic", "work"},
		},
		{
			Name:        "Keyboard",
			Type:        "Electronic",
			Description: "A mechanical keyboard",
			StorageID:   storage.ID,
			Tags:        []string{"electronic", "work"},
		},
		{
			Name:        "Mouse",
			Type:        "Electronic",
			Description: "A wireless mouse",
			StorageID:   storage.ID,
			Tags:        []string{"electronic", "work"},
		},
	}

	for _, item := range items {
		if err := database.CreateItem(context.Background(), &item); err != nil {
			log.Fatalf("failed to create item: %v", err)
		}
	}

	log.Println("Successfully seeded the database")
}
