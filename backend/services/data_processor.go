package services

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DataProcessor struct {
	db *gorm.DB
}

func NewDataProcessor(db *gorm.DB) *DataProcessor {
	return &DataProcessor{db: db}
}

// NFTBid represents the structure of an NFT bid event
type NFTBid struct {
	NFTAddress string  `json:"nft_address"`
	Bidder     string  `json:"bidder"`
	Amount     float64 `json:"amount"`
	Timestamp  int64   `json:"timestamp"`
}

// NFTPrice represents the structure of an NFT price event
type NFTPrice struct {
	NFTAddress string  `json:"nft_address"`
	Price      float64 `json:"price"`
	Market     string  `json:"market"`
	Timestamp  int64   `json:"timestamp"`
}

// TokenBorrow represents the structure of a token borrow event
type TokenBorrow struct {
	TokenAddress string  `json:"token_address"`
	Amount       float64 `json:"amount"`
	APY          float64 `json:"apy"`
	Platform     string  `json:"platform"`
	Timestamp    int64   `json:"timestamp"`
}

// TokenPrice represents the structure of a token price event
type TokenPrice struct {
	TokenAddress string  `json:"token_address"`
	Price        float64 `json:"price"`
	Platform     string  `json:"platform"`
	Timestamp    int64   `json:"timestamp"`
}

// ProcessNFTBid processes and stores NFT bid data
func (p *DataProcessor) ProcessNFTBid(event *models.WebhookEvent) error {
	var bid NFTBid
	if err := json.Unmarshal([]byte(event.Payload), &bid); err != nil {
		return fmt.Errorf("failed to parse NFT bid data: %v", err)
	}

	// Create a table dynamically for the user if it doesn't exist
	tableName := fmt.Sprintf("user_%d_nft_bids", event.WebhookID)
	p.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			nft_address TEXT NOT NULL,
			bidder TEXT NOT NULL,
			amount DECIMAL NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, tableName))

	// Insert the bid data
	sql := fmt.Sprintf(
		"INSERT INTO %s (nft_address, bidder, amount, timestamp) VALUES (?, ?, ?, ?)",
		tableName,
	)
	result := p.db.Exec(sql, bid.NFTAddress, bid.Bidder, bid.Amount, time.Unix(bid.Timestamp, 0))
	return result.Error
}

// ProcessNFTPrice processes and stores NFT price data
func (p *DataProcessor) ProcessNFTPrice(event *models.WebhookEvent) error {
	var price NFTPrice
	if err := json.Unmarshal([]byte(event.Payload), &price); err != nil {
		return fmt.Errorf("failed to parse NFT price data: %v", err)
	}

	tableName := fmt.Sprintf("user_%d_nft_prices", event.WebhookID)
	p.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			nft_address TEXT NOT NULL,
			price DECIMAL NOT NULL,
			market TEXT NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, tableName))

	sql := fmt.Sprintf(
		"INSERT INTO %s (nft_address, price, market, timestamp) VALUES (?, ?, ?, ?)",
		tableName,
	)
	result := p.db.Exec(sql, price.NFTAddress, price.Price, price.Market, time.Unix(price.Timestamp, 0))
	return result.Error
}

// ProcessTokenBorrow processes and stores token borrow data
func (p *DataProcessor) ProcessTokenBorrow(event *models.WebhookEvent) error {
	var borrow TokenBorrow
	if err := json.Unmarshal([]byte(event.Payload), &borrow); err != nil {
		return fmt.Errorf("failed to parse token borrow data: %v", err)
	}

	tableName := fmt.Sprintf("user_%d_token_borrows", event.WebhookID)
	p.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			token_address TEXT NOT NULL,
			amount DECIMAL NOT NULL,
			apy DECIMAL NOT NULL,
			platform TEXT NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, tableName))

	sql := fmt.Sprintf(
		"INSERT INTO %s (token_address, amount, apy, platform, timestamp) VALUES (?, ?, ?, ?, ?)",
		tableName,
	)
	result := p.db.Exec(sql, borrow.TokenAddress, borrow.Amount, borrow.APY, borrow.Platform, time.Unix(borrow.Timestamp, 0))
	return result.Error
}

// ProcessTokenPrice processes and stores token price data
func (p *DataProcessor) ProcessTokenPrice(event *models.WebhookEvent) error {
	var price TokenPrice
	if err := json.Unmarshal([]byte(event.Payload), &price); err != nil {
		return fmt.Errorf("failed to parse token price data: %v", err)
	}

	tableName := fmt.Sprintf("user_%d_token_prices", event.WebhookID)
	p.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			token_address TEXT NOT NULL,
			price DECIMAL NOT NULL,
			platform TEXT NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`, tableName))

	sql := fmt.Sprintf(
		"INSERT INTO %s (token_address, price, platform, timestamp) VALUES (?, ?, ?, ?)",
		tableName,
	)
	result := p.db.Exec(sql, price.TokenAddress, price.Price, price.Platform, time.Unix(price.Timestamp, 0))
	return result.Error
}
