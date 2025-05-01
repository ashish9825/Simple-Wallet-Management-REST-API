package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ashish9825/wallet-api/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Read the DB path from environment variable
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "wallet.db" // default value if not set
	}

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite DB:", err)
	}

	fmt.Printf("Connected to SQLite DB at %s\n", dbPath)

	DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
}

