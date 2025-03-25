package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
    gorm.Model
    Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	ApiKey       string `json:"api_key" gorm:"unique"`
    DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	DbName      string `json:"db_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type IndexingPreference struct {
    gorm.Model
    UserID           uint	 
    NFTBids         bool
    NFTPrices       bool
    BorrowableTokens bool
    TokenPrices     bool
    CustomFilters   string `gorm:"type:json" json:"custom_filters"` // Changed to string with json type
}

type DataSyncStatus struct {
    gorm.Model
    UserID         uint
    LastSynced     time.Time `gorm:"type:timestamp"`
    SyncedBlocks   int
    ErrorLog       string
}
