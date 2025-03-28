package main

import (
"backend/config"
	"backend/controllers"
	"backend/middleware"
	"backend/models"
	"backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Database connection
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the schema
	// Auto-migrate all models
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

	// Initialize Fiber
	app := fiber.New()

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Blockchain Indexing Platform API")
	})

	// Auth routes
	app.Post("/auth/signup", userController.Signup)
	app.Post("/auth/login", userController.Login)

	// Protected webhook routes
	app.Post("/webhooks/helius", middleware.Protected(), webhookController.HandleWebhook)
	app.Post("/webhooks/configure", middleware.Protected(), webhookController.ConfigureWebhook)

	// Start server
	app.Listen(":" + cfg.ServerPort)
}