package models

import (
	"time"

	"gorm.io/gorm"
)

// HeliusWebhook represents the webhook configuration for a user
type HeliusWebhook struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	WebhookURL  string `json:"webhook_url" gorm:"unique"`
	WebhookID   string `json:"webhook_id" gorm:"unique"`
	AccountKeys string `json:"account_keys" gorm:"type:text[]"` // Array of account addresses to monitor
	EventTypes  string `json:"event_types" gorm:"type:text[]"`  // Array of event types to monitor
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

// WebhookEvent represents an event received from Helius
type WebhookEvent struct {
	gorm.Model
	WebhookID    uint      `json:"webhook_id"`
	EventType    string    `json:"event_type"`
	AccountKey   string    `json:"account_key"`
	Slot         uint64    `json:"slot"`
	Timestamp    time.Time `json:"timestamp"`
	Payload      string    `json:"payload" gorm:"type:jsonb"`
	Processed    bool      `json:"processed" gorm:"default:false"`
	ProcessedAt  time.Time `json:"processed_at"`
	ErrorMessage string    `json:"error_message"`
}
