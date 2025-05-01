package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ashish9825/wallet-api/database"
	"github.com/ashish9825/wallet-api/routes"
	
)

func main() {
	database.ConnectDB()
	r := gin.Default()


	routes.SetupRoutes(r)
	r.Run(":8080")
}