package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// JWT secret key (should be in .env in production)
var jwtKey = []byte("your_secret_jwt_key")

// Model structs
type Store struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	WhatsappNumber string    `json:"whatsappNumber"`
	Logo           string    `json:"logo"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type Product struct {
	ID          string    `json:"id"`
	StoreID     string    `json:"storeId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	StoreID  string `json:"storeId"`
}

// In-memory storage
var stores = []Store{}
var products = []Product{}
var users = []User{}

// Generate JWT token
func generateToken(username string, storeID string) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims
	claims := jwt.MapClaims{
		"username": username,
		"storeId":  storeID,
		"exp":      expirationTime.Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWT Middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Missing authorization token"})
			return
		}

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
			return
		}

		// Token is valid, proceed
		next.ServeHTTP(w, r)
	})
}

// Handler functions
func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Check if username already exists
	for _, existingUser := range users {
		if existingUser.Username == user.Username {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Username already exists"})
			return
		}
	}

	// Generate unique ID
	user.ID = fmt.Sprintf("user_%d", time.Now().UnixNano())

	// Store user
	users = append(users, user)

	// Generate token
	token, err := generateToken(user.Username, user.StoreID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
		return
	}

	// Return token
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData User
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Find user
	var user User
	found := false
	for _, u := range users {
		if u.Username == loginData.Username && u.Password == loginData.Password {
			user = u
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := generateToken(user.Username, user.StoreID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
		return
	}

	// Return token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func createStoreHandler(w http.ResponseWriter, r *http.Request) {
	var store Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Generate unique ID
	store.ID = fmt.Sprintf("store_%d", time.Now().UnixNano())
	
	// Set timestamps
	store.CreatedAt = time.Now()
	store.UpdatedAt = time.Now()

	// Store data
	stores = append(stores, store)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(store)
}

func getStoreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID := vars["id"]

	// Find store
	var store Store
	found := false
	for _, s := range stores {
		if s.ID == storeID {
			store = s
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Store not found"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(store)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID := vars["storeId"]

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Generate unique ID
	product.ID = fmt.Sprintf("product_%d", time.Now().UnixNano())
	product.StoreID = storeID
	
	// Set timestamps
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	// Store data
	products = append(products, product)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID := vars["storeId"]

	// Filter products by store ID
	var storeProducts []Product
	for _, product := range products {
		if product.StoreID == storeID {
			storeProducts = append(storeProducts, product)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storeProducts)
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var updatedProduct Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	// Find and update product
	found := false
	for i, product := range products {
		if product.ID == productID {
			// Update fields
			products[i].Name = updatedProduct.Name
			products[i].Description = updatedProduct.Description
			products[i].Price = updatedProduct.Price
			products[i].Category = updatedProduct.Category
			products[i].Image = updatedProduct.Image
			products[i].UpdatedAt = time.Now()
			found = true
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
	}
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	// Find and delete product
	for i, product := range products {
		if product.ID == productID {
			// Remove product
			products = append(products[:i], products[i+1:]...)
			
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
}

func main() {
	// Load .env file if it exists
	godotenv.Load()

	// Router setup
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/api/auth/register", registerHandler).Methods("POST")
	r.HandleFunc("/api/auth/login", loginHandler).Methods("POST")
	r.HandleFunc("/api/stores/{id}", getStoreHandler).Methods("GET")
	r.HandleFunc("/api/stores/{storeId}/products", getProductsHandler).Methods("GET")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(authMiddleware)
	api.HandleFunc("/stores", createStoreHandler).Methods("POST")
	api.HandleFunc("/stores/{storeId}/products", createProductHandler).Methods("POST")
	api.HandleFunc("/products/{id}", updateProductHandler).Methods("PUT")
	api.HandleFunc("/products/{id}", deleteProductHandler).Methods("DELETE")

	// CORS setup
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Start server
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(r)))
}