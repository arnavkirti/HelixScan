package services

import (
	"backend/models"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type HeliusService struct {
	db            *gorm.DB
	heliusApiKey  string
	dataProcessor *DataProcessor
	apiClient     *HeliusAPIClient
}

func NewHeliusService(db *gorm.DB, apiKey string) *HeliusService {
	return &HeliusService{
		db:            db,
		heliusApiKey:  apiKey,
		dataProcessor: NewDataProcessor(db),
		apiClient:     NewHeliusAPIClient(apiKey),
	}
}

// RegisterWebhook creates a new webhook configuration for a user
func (s *HeliusService) RegisterWebhook(userID uint, accountKeys []string, eventTypes []string) (*models.HeliusWebhook, error) {
	// Create webhook configuration in Helius
	webhookURL := fmt.Sprintf("%s/webhooks/helius", "https://your-api-domain.com") // Replace with your actual domain

	// Register webhook with Helius API
	resp, err := s.apiClient.RegisterWebhook(WebhookRegistrationRequest{
		WebhookURL:       webhookURL,
		AccountAddresses: accountKeys,
		EventTypes:       eventTypes,
		AuthHeader:       "Bearer " + s.heliusApiKey, // Add auth header for webhook security
	})

	if err != nil {
		return nil, fmt.Errorf("failed to register webhook with Helius: %w", err)
	}

	// Create webhook record in database
	webhook := &models.HeliusWebhook{
		UserID:      userID,
		WebhookURL:  webhookURL,
		WebhookID:   resp.WebhookID,
		AccountKeys: strings.Join(accountKeys, ","), // Store as comma-separated string
		EventTypes:  strings.Join(eventTypes, ","),  // Store as comma-separated string
		IsActive:    true,
	}

	result := s.db.Create(webhook)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to store webhook in database: %w", result.Error)
	}

	return webhook, nil
}

// ProcessWebhookEvent handles incoming webhook events from Helius
func (s *HeliusService) ProcessWebhookEvent(event *models.WebhookEvent) error {
	// Validate webhook exists and is active
	var webhook models.HeliusWebhook
	result := s.db.First(&webhook, event.WebhookID)
	if result.Error != nil {
		return errors.New("webhook not found")
	}

	if !webhook.IsActive {
		return errors.New("webhook is inactive")
	}

	// Store the event
	event.Timestamp = time.Now()
	result = s.db.Create(event)
	if result.Error != nil {
		return result.Error
	}

	// Process the event based on type
	switch event.EventType {
	case "nft_bid":
		return s.processNFTBid(event)
	case "nft_price":
		return s.processNFTPrice(event)
	case "token_borrow":
		return s.processTokenBorrow(event)
	case "token_price":
		return s.processTokenPrice(event)
	default:
		return fmt.Errorf("unsupported event type: %s", event.EventType)
	}
}

// Helper functions for processing different event types
func (s *HeliusService) processNFTBid(event *models.WebhookEvent) error {
	return s.dataProcessor.ProcessNFTBid(event)
}

func (s *HeliusService) processNFTPrice(event *models.WebhookEvent) error {
	return s.dataProcessor.ProcessNFTPrice(event)
}

func (s *HeliusService) processTokenBorrow(event *models.WebhookEvent) error {
	return s.dataProcessor.ProcessTokenBorrow(event)
}

func (s *HeliusService) processTokenPrice(event *models.WebhookEvent) error {
	return s.dataProcessor.ProcessTokenPrice(event)
}
