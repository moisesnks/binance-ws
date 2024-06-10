package handlers

import (
	"backend-ws/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUser crea un nuevo usuario
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.Usuario
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Crear una billetera para el nuevo usuario
		newWallet := models.Billetera{UsuarioID: newUser.UID}
		if err := db.Create(&newWallet).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newUser)
	}
}

// GetUsers obtiene todos los usuarios
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.Usuario
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// CreateWallet crea una nueva billetera
func CreateWallet(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newWallet models.Billetera
		if err := c.ShouldBindJSON(&newWallet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newWallet).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newWallet)
	}
}

// GetWallets obtiene todas las billeteras
func GetWallets(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var wallets []models.Billetera
		if err := db.Find(&wallets).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, wallets)
	}
}

// CreateCurrency crea una nueva moneda
func CreateCurrency(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCurrency models.Moneda
		if err := c.ShouldBindJSON(&newCurrency); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newCurrency).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newCurrency)
	}
}

// GetCurrencies obtiene todas las monedas
func GetCurrencies(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var currencies []models.Moneda
		if err := db.Find(&currencies).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, currencies)
	}
}

// AddCoins agrega monedas a una billetera existente
func AddCoins(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			UserID   string `json:"user_id"`
			MonedaID uint   `json:"moneda_id"`
			Cantidad int    `json:"cantidad"`
		}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Buscar la billetera asociada al usuario
		var billetera models.Billetera
		if err := db.Where("usuario_id = ?", requestData.UserID).First(&billetera).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Verificar si la moneda ya está asociada a la billetera
		var moneda models.Moneda
		if err := db.First(&moneda, requestData.MonedaID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&billetera).Association("Monedas").Append(&moneda); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Actualizar la cantidad de monedas en la billetera
		moneda.Cantidad += requestData.Cantidad
		if err := db.Save(&moneda).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Monedas agregadas exitosamente"})
	}
}
