// controllers/shipping_controller.go
package shipping_controller

import (
	"atlanta-site/models"
	services "atlanta-site/services/shipping_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetShippingOptions retorna as opções de envio disponíveis via LOGGI
// @Summary Obtém opções de envio disponíveis
// @Description Retorna uma lista de métodos de entrega e seus preços
// @Tags Shipping
// @Produce json
// @Success 200 {array} models.ShippingOption
// @Failure 500 {object} map[string]string "Erro ao buscar opções de envio"
// @Router /shipping/options [get]
func GetShippingOptions(c *gin.Context) {
	options, err := services.GetShippingOptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar opções de envio"})
		return
	}
	c.JSON(http.StatusOK, options)
}

// CreateShipping cria um pedido de envio via LOGGI
// @Summary Cria um pedido de envio
// @Description Registra um novo pedido de envio e retorna os detalhes do tracking
// @Tags Shipping
// @Accept json
// @Produce json
// @Param request body models.ShippingRequest true "Detalhes do envio"
// @Success 200 {object} models.ShippingResponse
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao criar envio"
// @Router /shipping/create [post]
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
