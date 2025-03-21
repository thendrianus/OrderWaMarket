package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Collection names
const (
	UserCollection    = "users"
	StoreCollection   = "stores"
	ProductCollection = "products"
)

// User represents a user document in MongoDB
type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Username     string             `bson:"username" json:"username"`
	PasswordHash string             `bson:"password_hash" json:"-"` // Not included in JSON responses
	Email        string             `bson:"email" json:"email"`
	Role         string             `bson:"role" json:"role"` // "admin", "owner", etc.
	StoreID      primitive.ObjectID `bson:"store_id,omitempty" json:"storeId,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updatedAt"`
}

// Store represents a store document in MongoDB
type Store struct {
	ID             primitive.ObjectID  `bson:"_id" json:"id"`
	OwnerID        primitive.ObjectID  `bson:"owner_id" json:"ownerId"`
	Name           string              `bson:"name" json:"name"`
	Description    string              `bson:"description" json:"description"`
	Logo           string              `bson:"logo" json:"logo"`
	Location       string              `bson:"location" json:"location"`
	WhatsappNumber string              `bson:"whatsapp_number" json:"whatsappNumber"`
	BusinessHours  string              `bson:"business_hours" json:"businessHours"`
	Tags           []string            `bson:"tags,omitempty" json:"tags,omitempty"`
	Active         bool                `bson:"active" json:"active"`
	FeaturedProduct primitive.ObjectID `bson:"featured_product,omitempty" json:"featuredProduct,omitempty"`
	CreatedAt      time.Time           `bson:"created_at" json:"createdAt"`
	UpdatedAt      time.Time           `bson:"updated_at" json:"updatedAt"`
}

// Product represents a product document in MongoDB
type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	StoreID     primitive.ObjectID `bson:"store_id" json:"storeId"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Image       string             `bson:"image" json:"image"`
	Category    string             `bson:"category" json:"category"`
	Stock       int                `bson:"stock" json:"stock"`
	Featured    bool               `bson:"featured" json:"featured"`
	Active      bool               `bson:"active" json:"active"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}

// API Request/Response Models

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse represents the response body for auth endpoints
type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateStoreRequest represents the request body for store creation
type CreateStoreRequest struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Logo           string   `json:"logo"`
	Location       string   `json:"location"`
	WhatsappNumber string   `json:"whatsappNumber"`
	BusinessHours  string   `json:"businessHours"`
	Tags           []string `json:"tags,omitempty"`
}

// UpdateStoreRequest represents the request body for store updates
type UpdateStoreRequest struct {
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Logo           string   `json:"logo,omitempty"`
	Location       string   `json:"location,omitempty"`
	WhatsappNumber string   `json:"whatsappNumber,omitempty"`
	BusinessHours  string   `json:"businessHours,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	Active         *bool    `json:"active,omitempty"`
}

// CreateProductRequest represents the request body for product creation
type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Category    string   `json:"category"`
	Stock       int      `json:"stock"`
	Featured    bool     `json:"featured"`
	Active      *bool    `json:"active,omitempty"`
}

// UpdateProductRequest represents the request body for product updates
type UpdateProductRequest struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Image       string  `json:"image,omitempty"`
	Category    string  `json:"category,omitempty"`
	Stock       *int    `json:"stock,omitempty"`
	Featured    *bool   `json:"featured,omitempty"`
	Active      *bool   `json:"active,omitempty"`
}