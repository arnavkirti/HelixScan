package controllers

import (
	"backend/models"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

type WebhookController struct {
	heliusService *services.HeliusService
}

func NewWebhookController(heliusService *services.HeliusService) *WebhookController {
	return &WebhookController{
		heliusService: heliusService,
	}
}

// HandleWebhook processes incoming webhook events from Helius
func (c *WebhookController) HandleWebhook(ctx *fiber.Ctx) error {
	var event models.WebhookEvent
	if err := ctx.BodyParser(&event); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.heliusService.ProcessWebhookEvent(&event); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event processed successfully",
	})
}

// ConfigureWebhook allows users to set up their webhook preferences
func (c *WebhookController) ConfigureWebhook(ctx *fiber.Ctx) error {
	type WebhookConfig struct {
		UserID      uint     `json:"user_id"`
		AccountKeys []string `json:"account_keys"`
		EventTypes  []string `json:"event_types"`
	}

	var config WebhookConfig
	if err := ctx.BodyParser(&config); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	webhook, err := c.heliusService.RegisterWebhook(config.UserID, config.AccountKeys, config.EventTypes)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(webhook)
}