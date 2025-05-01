package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ashish9825/wallet-api/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.POST("/wallets", controllers.CreateWallet)
	router.GET("/wallets/:id/balance", controllers.GetBalance)
	router.POST("/wallets/transfer", controllers.TransferFunds)
	router.GET("/wallets/:id/transactions", controllers.GetTransactions)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/wallets", controllers.GetAllWallets)
	router.GET("/transactions", controllers.GetAllTransactions)

}
