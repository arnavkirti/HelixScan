package main

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func main() {
	// Database connection
	dsn := "host=localhost user=bounty password=bounty123 dbname=bounty_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the schema
	db.AutoMigrate(&models.User{}, &models.DatabaseConfig{}, &models.IndexingPreference{})

	// Initialize Fiber
	app := fiber.New()

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Blockchain Indexing Platform API")
	})

	// Start server
	app.Listen(":3000")
}
