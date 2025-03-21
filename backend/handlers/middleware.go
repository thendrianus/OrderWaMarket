package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthMiddleware checks for a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// Expected format: "Bearer {token}"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			RespondWithError(w, http.StatusUnauthorized, "Invalid authorization format, expected 'Bearer {token}'")
			return
		}

		tokenString := tokenParts[1]

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("Invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Add user ID and username to request context
			userID, err := primitive.ObjectIDFromHex(claims["userId"].(string))
			if err != nil {
				RespondWithError(w, http.StatusUnauthorized, "Invalid user ID in token")
				return
			}

			username := claims["username"].(string)

			// Create a new context with user info
			ctx := context.WithValue(r.Context(), "userID", userID)
			ctx = context.WithValue(ctx, "username", username)

			// Call next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			RespondWithError(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}
	})
}

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

// Helper function to get user ID from request context
func getUserIDFromContext(r *http.Request) (primitive.ObjectID, error) {
	userID, ok := r.Context().Value("userID").(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, jwt.NewValidationError("User ID not found in context", jwt.ValidationErrorMalformed)
	}
	return userID, nil
}

// Helper function to get username from request context
func getUsernameFromContext(r *http.Request) (string, error) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		return "", jwt.NewValidationError("Username not found in context", jwt.ValidationErrorMalformed)
	}
	return username, nil
}