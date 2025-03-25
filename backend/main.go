package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	// Connect to PostgreSQL
	connStr := "user=bounty dbname=bounty_db password=bounty123 host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/signup", SignupHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", LoginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/save-db-config", SaveDBConfigHandler).Methods("POST", "OPTIONS")

	// Add CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			enableCORS(&w)
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Insert into database
	var userID int
	err = db.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", user.Email, string(hashedPassword)).Scan(&userID)
	if err != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User created", "user_id": userID})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Get user from database
	var dbUser struct {
		ID           int
		Email        string
		PasswordHash string
	}
	err = db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", user.Email).Scan(&dbUser.ID, &dbUser.Email, &dbUser.PasswordHash)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful", "user_id": dbUser.ID})
}

func SaveDBConfigHandler(w http.ResponseWriter, r *http.Request) {
    var config struct {
        Host         string `json:"host"`
        Port         int    `json:"port"`
        DatabaseName string `json:"db_name"`
        Username     string `json:"db_username"`
        Password     string `json:"db_password"`
        UserID       int    `json:"user_id"`
    }
    
    err := json.NewDecoder(r.Body).Decode(&config)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    _, err = db.Exec(`
        INSERT INTO user_databases 
        (user_id, host, port, db_name, db_username, db_password) 
        VALUES ($1, $2, $3, $4, $5, $6)`,
        config.UserID, config.Host, config.Port, config.DatabaseName, config.Username, config.Password)
    if err != nil {
        http.Error(w, "Failed to save config", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Database config saved"})
}