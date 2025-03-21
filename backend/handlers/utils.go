package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// JWT secret key - in production, this should be stored in environment variable
var jwtSecret = os.Getenv("JWT_SECRET")

// Initialize JWT secret if not set in environment
func init() {
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-for-whatsapp-catalogue" // Default fallback for development
	}
}

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// RespondWithError sends an error response
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, ErrorResponse{Status: code, Message: message})
}

// RespondWithJSON sends a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":500,"message":"Internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a bcrypt hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID primitive.ObjectID, username string) (string, error) {
	// Set token expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims
	claims := jwt.MapClaims{
		"userId":   userID.Hex(),
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}