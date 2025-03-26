package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
    gorm.Model
    Email              string `json:"email" gorm:"unique;not null"`
    Password           string `json:"password" gorm:"not null"`
    ApiKey             string `json:"api_key" gorm:"unique;not null"`
    DBConfig           DatabaseConfig     `gorm:"foreignKey:UserID"`
    IndexingPreference IndexingPreference `gorm:"foreignKey:UserID"`
    SyncStatus         []DataSyncStatus   `gorm:"foreignKey:UserID"`
}

type DatabaseConfig struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"not null;uniqueIndex"`
	Host        string `json:"host" gorm:"not null"`
	Port        string `json:"port" gorm:"not null"`
	DbName      string `json:"db_name" gorm:"not null"`
	Username    string `json:"username" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
}

type IndexingPreference struct {
	gorm.Model
	UserID          uint   `json:"user_id" gorm:"not null;uniqueIndex"`
	NFTBids         bool   `json:"nft_bids" gorm:"default:true"`
	NFTPrices       bool   `json:"nft_prices" gorm:"default:true"`
	BorrowableTokens bool  `json:"borrowable_tokens" gorm:"default:false"`
	TokenPrices     bool   `json:"token_prices" gorm:"default:false"`
	CustomFilters   string `json:"custom_filters" gorm:"type:json"`
}

type DataSyncStatus struct {
	gorm.Model
	UserID         uint      `json:"user_id" gorm:"not null;uniqueIndex"`
	LastSynced     time.Time `json:"last_synced" gorm:"type:timestamp"`
	SyncedBlocks   int       `json:"synced_blocks" gorm:"default:0"`
	ErrorLog       string    `json:"error_log"`
}
