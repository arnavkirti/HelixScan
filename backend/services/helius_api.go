package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HeliusAPIClient struct {
	apiKey  string
	baseURL string
}

type WebhookRegistrationRequest struct {
	WebhookURL       string   `json:"webhookURL"`
	AccountAddresses []string `json:"accountAddresses"`
	EventTypes       []string `json:"eventTypes"`
	AuthHeader       string   `json:"authHeader,omitempty"`
}

type WebhookRegistrationResponse struct {
	WebhookID string `json:"webhookID"`
}

func NewHeliusAPIClient(apiKey string) *HeliusAPIClient {
	return &HeliusAPIClient{
		apiKey:  apiKey,
		baseURL: "https://api.helius.xyz/v0",
	}
}

// RegisterWebhook registers a new webhook with Helius
func (c *HeliusAPIClient) RegisterWebhook(req WebhookRegistrationRequest) (*WebhookRegistrationResponse, error) {
	url := fmt.Sprintf("%s/webhooks", c.baseURL)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response WebhookRegistrationResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
