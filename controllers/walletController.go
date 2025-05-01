package controllers

import (
	"net/http"
	"time"
    "gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/ashish9825/wallet-api/database"
	"github.com/ashish9825/wallet-api/models"
)


func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func CreateWallet(c *gin.Context) {
	var wallet models.Wallet
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	database.DB.Create(&wallet)
	c.JSON(http.StatusOK, wallet)
}

func GetBalance(c *gin.Context) {
	var wallet models.Wallet
	id := c.Param("id")
	if err := database.DB.First(&wallet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": wallet.Balance})
}

func TransferFunds(c *gin.Context) {
	var req struct {
		FromID uint    `json:"from_wallet_id"`
		ToID   uint    `json:"to_wallet_id"`
		Amount float64 `json:"amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var from, to models.Wallet
	if err := database.DB.First(&from, req.FromID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sender wallet not found"})
		return
	}
	if from.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}
	if err := database.DB.First(&to, req.ToID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receiver wallet not found"})
		return
	}

	database.DB.Transaction(func(tx *gorm.DB) error {
		from.Balance -= req.Amount
		to.Balance += req.Amount

		tx.Save(&from)
		tx.Save(&to)

		tx.Create(&models.Transaction{
			FromWalletID: req.FromID,
			ToWalletID:   req.ToID,
			Amount:       req.Amount,
			Timestamp:    time.Now(),
		})

		return nil
	})

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	id := c.Param("id")
	database.DB.
		Where("from_wallet_id = ? OR to_wallet_id = ?", id, id).
		Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(200, users)
}

func GetAllWallets(c *gin.Context) {
	var wallets []models.Wallet
	database.DB.Find(&wallets)
	c.JSON(200, wallets)
}

func GetAllTransactions(c *gin.Context) {
	var txs []models.Transaction
	database.DB.Find(&txs)
	c.JSON(200, txs)
}

