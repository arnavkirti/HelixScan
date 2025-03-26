package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"

	"backend/models" // Update import path to match your module name
)

var db *gorm.DB

func main() {
	// Initialize database connection
	dsn := "host=localhost user=bounty password=bounty123 dbname=bounty_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.DatabaseConfig{},
		&models.IndexingPreference{},
		&models.DataSyncStatus{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize router with CORS middleware
	r := mux.NewRouter()
	r.Use(corsMiddleware)

	// Routes
	r.HandleFunc("/signup", signupHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", loginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/db-config", dbConfigHandler).Methods("GET", "POST", "OPTIONS")

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Add the missing dbConfigHandler
func dbConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		var configs []models.DatabaseConfig
		if err := db.Find(&configs).Error; err != nil {
			http.Error(w, "Failed to fetch configs", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(configs)

	case "POST":
		var config models.DatabaseConfig
		if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := db.Create(&config).Error; err != nil {
			http.Error(w, "Failed to save config", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(config)
	}
}

// Update signupHandler to return user ID and API key
func signupHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    var request struct {
        Email    string             `json:"email"`
        Password string             `json:"password"`
        Database models.DatabaseConfig `json:"database"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Generate API key
    apiKey := generateAPIKey()

    var user models.User
    // Create user in transaction
    err = db.Transaction(func(tx *gorm.DB) error {
        user = models.User{
            Email:    request.Email,
            Password: string(hashedPassword),
            ApiKey:   apiKey,
        }

        if err := tx.Create(&user).Error; err != nil {
            return err
        }

        // Create database config
        dbConfig := models.DatabaseConfig{
            UserID:   user.ID,
            Host:     request.Database.Host,
            Port:     request.Database.Port,
            DbName:   request.Database.DbName,
            Username: request.Database.Username,
            Password: request.Database.Password,
        }
        if err := tx.Create(&dbConfig).Error; err != nil {
            return err
        }

        // Create indexing preferences
        prefs := models.IndexingPreference{
            UserID:           user.ID,
            NFTBids:         true,
            NFTPrices:       true,
            BorrowableTokens: false,
            TokenPrices:     false,
        }
        return tx.Create(&prefs).Error
    })

    if err != nil {
        http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User created successfully",
        "user_id": user.ID,
        "api_key": user.ApiKey,
    })
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return user data without password
	response := struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		ApiKey   string `json:"api_key"`
	}{
		ID:       user.ID,
		Email:    user.Email,
		ApiKey:   user.ApiKey,
	}

	json.NewEncoder(w).Encode(response)
}

func generateAPIKey() string {
	return "sk_" + time.Now().Format("20060102150405") + "_" + fmt.Sprintf("%x", time.Now().UnixNano())
}