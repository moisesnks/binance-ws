package main

import (
	"backend-ws/config"
	"backend-ws/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Definir las rutas y los handlers
	r.POST("/users", handlers.CreateUser(db))
	r.GET("/users", handlers.GetUsers(db))
	r.POST("/wallets", handlers.CreateWallet(db))
	r.GET("/wallets", handlers.GetWallets(db))
	r.POST("/currencies", handlers.CreateCurrency(db))
	r.GET("/currencies", handlers.GetCurrencies(db))
	r.POST("/add-coins", handlers.AddCoins(db))

	// Iniciar el servidor
	r.Run(":8080")
}
