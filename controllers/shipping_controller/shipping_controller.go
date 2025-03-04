// controllers/shipping_controller.go
package shipping_controller

import (
	"atlanta-site/models"
	services "atlanta-site/services/shipping_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetShippingOptions obtém opções de envio disponíveis via LOGGI
func GetShippingOptions(c *gin.Context) {
	options, err := services.GetShippingOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar opções de envio"})
		return
	}
	c.JSON(http.StatusOK, options)
}

// CreateShipping cria um pedido de envio via LOGGI
func CreateShipping(c *gin.Context) {
	var shippingRequest models.ShippingRequest
	if err := c.ShouldBindJSON(&shippingRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	shippingResponse, err := services.CreateShipping(shippingRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar envio"})
		return
	}

	c.JSON(http.StatusOK, shippingResponse)
}
