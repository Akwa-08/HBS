package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing tokens (Use ENV variables in production)
var jwtSecret = []byte("this_is_a_much_longer_super_secret_key_123!")

// Database connection
var db *sql.DB

// User struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

// Initialize DB
func initDB() {
	var err error
	connStr := "postgres://postgres:password@auth_db:5432/auth?sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func CorsMiddleware() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend's origin
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	return c
}

// Generate JWT token
func generateToken(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":                    userID,
		"exp":                    time.Now().Add(time.Hour * 24).Unix(),
		"x-hasura-allowed-roles": []string{"user", "admin"},
		"x-hasura-default-role":  role,
		"x-hasura-user-id":       userID,
		"x-hasura-role":          role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Insert user into DB and retrieve the generated ID
    var userID int
    err = db.QueryRow(
        "INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id",
        user.Username, string(hashedPassword), user.Role,
    ).Scan(&userID)
    if err != nil {
        http.Error(w, "Error saving user", http.StatusInternalServerError)
        return
    }

    // Generate JWT
    token, err := generateToken(fmt.Sprintf("%d", userID), user.Role)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    // Return token and user details
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "token": token,
        "user": map[string]interface{}{
            "id":   userID,
            "role": user.Role,
        },
    })

	log.Println("Generated user ID:", userID)
}

// Login user
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Retrieve user from DB
	var storedUser User
	err = db.QueryRow("SELECT id, password, role FROM users WHERE username = $1", user.Username).Scan(&storedUser.ID, &storedUser.Password, &storedUser.Role)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := generateToken(fmt.Sprintf("%d", storedUser.ID), storedUser.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Return token and user details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":   storedUser.ID,
			"role": storedUser.Role,
		},
	})
}

// Token validation endpoint
func validateTokenHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " prefix if present
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Return decoded claims
	claims, _ := token.Claims.(jwt.MapClaims)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
}

// Change password handler
func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the request body
    var req struct {
        OldPassword string `json:"old_password"`
        NewPassword string `json:"new_password"`
    }
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Get the token from the Authorization header
    tokenStr := r.Header.Get("Authorization")
    if tokenStr == "" {
        http.Error(w, "Missing token", http.StatusUnauthorized)
        return
    }

    // Remove "Bearer " prefix if present
    if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
        tokenStr = tokenStr[7:]
    }

    // Parse and validate the token
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil || !token.Valid {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }

    // Extract user ID from the token claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        http.Error(w, "Invalid token claims", http.StatusUnauthorized)
        return
    }
    userID := claims["x-hasura-user-id"].(string)

    // Retrieve the user's current password from the database
    var storedPassword string
    err = db.QueryRow("SELECT password FROM users WHERE id = $1", userID).Scan(&storedPassword)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Verify the old password
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(req.OldPassword))
    if err != nil {
        http.Error(w, "Old password is incorrect", http.StatusUnauthorized)
        return
    }

    // Hash the new password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Update the password in the database
    _, err = db.Exec("UPDATE users SET password = $1 WHERE id = $2", string(hashedPassword), userID)
    if err != nil {
        http.Error(w, "Error updating password", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}

// Main function
func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateTokenHandler)
	http.HandleFunc("/change-password", changePasswordHandler) 

	corsHandler := CorsMiddleware().Handler(http.DefaultServeMux)
	log.Println("Auth Service running on :4000")
	log.Fatal(http.ListenAndServe(":4000", corsHandler))
}
