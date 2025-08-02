package models

import (
	"time"

	"github.com/google/uuid"
)

// Storage represents a storage container that can contain items and other storages
type Storage struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Location    string     `json:"location" db:"location"`
	Description string     `json:"description" db:"description"`
	ImageURL    *string    `json:"image_url" db:"image_url"`
	ParentID    *uuid.UUID `json:"parent_id" db:"parent_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// Item represents a physical item in the inventory
type Item struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"`
	Description string    `json:"description" db:"description"`
	ImageURL    *string   `json:"image_url" db:"image_url"`
	StorageID   uuid.UUID `json:"storage_id" db:"storage_id"`
	Tags        []string  `json:"tags" db:"tags"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Tag represents a label/tag that can be applied to items
type Tag struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

// QRCode represents generated QR codes for items and storages
type QRCode struct {
	ID         uuid.UUID `json:"id" db:"id"`
	EntityID   uuid.UUID `json:"entity_id" db:"entity_id"`     // ID of item or storage
	EntityType string    `json:"entity_type" db:"entity_type"` // "item" or "storage"
	CodeData   string    `json:"code_data" db:"code_data"`     // QR code content
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
