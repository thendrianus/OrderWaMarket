package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"wacatalogue/backend/handlers"
	"wacatalogue/backend/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Connect to database
	db, err := models.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create router
	router := mux.NewRouter()

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Public routes
	apiRouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		handlers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
	}).Methods("GET")

	// Auth routes
	apiRouter.HandleFunc("/auth/register", handlers.Register(db)).Methods("POST")
	apiRouter.HandleFunc("/auth/login", handlers.Login(db)).Methods("POST")

	// Store routes (public)
	apiRouter.HandleFunc("/stores", handlers.GetAllStores(db)).Methods("GET")
	apiRouter.HandleFunc("/stores/{id}", handlers.GetStore(db)).Methods("GET")
	apiRouter.HandleFunc("/stores/{storeId}/products", handlers.GetStoreProducts(db)).Methods("GET")
	apiRouter.HandleFunc("/products/{id}", handlers.GetProduct(db)).Methods("GET")

	// Protected routes
	protectedRouter := apiRouter.PathPrefix("/").Subrouter()
	protectedRouter.Use(handlers.AuthMiddleware)

	// Store routes (protected)
	protectedRouter.HandleFunc("/my-store", handlers.GetMyStore(db)).Methods("GET")
	protectedRouter.HandleFunc("/stores", handlers.CreateStore(db)).Methods("POST")
	protectedRouter.HandleFunc("/stores/{id}", handlers.UpdateStore(db)).Methods("PUT")
	protectedRouter.HandleFunc("/stores/{id}", handlers.DeleteStore(db)).Methods("DELETE")

	// Product routes (protected)
	protectedRouter.HandleFunc("/stores/{storeId}/products", handlers.CreateProduct(db)).Methods("POST")
	protectedRouter.HandleFunc("/products/{id}", handlers.UpdateProduct(db)).Methods("PUT")
	protectedRouter.HandleFunc("/products/{id}", handlers.DeleteProduct(db)).Methods("DELETE")

	// CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Create server
	handler := c.Handler(router)
	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,
	}

	// Start server
	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}