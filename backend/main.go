package main

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"backend/models"
	"backend/services"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL database
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate database models
	err = db.AutoMigrate(
		&models.User{},
		&models.DatabaseConfig{},
		&models.IndexingPreference{},
		&models.DataSyncStatus{},
		&models.HeliusWebhook{},
		&models.WebhookEvent{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize services
	heliusService := services.NewHeliusService(db, cfg.HeliusAPIKey)

	// Initialize controllers
	webhookController := controllers.NewWebhookController(heliusService)
	userController := controllers.NewUserController(db)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	// CORS Middleware (Allow requests from frontend)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",  // Frontend URL
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
		ExposeHeaders:    "Authorization",
		AllowCredentials: true,
	}))

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Blockchain Indexing Platform API")
	})

	// Authentication routes
	app.Post("/auth/signup", userController.Signup)
	app.Post("/auth/login", userController.Login)

	// Protected webhook routes
	app.Post("/webhooks/helius", middleware.Protected(), webhookController.HandleWebhook)
	app.Post("/webhooks/configure", middleware.Protected(), webhookController.ConfigureWebhook)

	// Start server on configured port
	port := "3001" // Changed to 3001 to avoid conflict with frontend
	
	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
