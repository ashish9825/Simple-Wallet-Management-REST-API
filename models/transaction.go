package models

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	FromWalletID  uint      `json:"from_wallet_id"`
	ToWalletID    uint      `json:"to_wallet_id"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}
