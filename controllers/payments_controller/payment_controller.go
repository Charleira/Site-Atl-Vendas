// controllers/payment_controller.go
package payment_controller

import (
	"atlanta-site/models"
	services "atlanta-site/services/stripe_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePaymentIntent inicia o processo de pagamento via Stripe Checkout
// @Summary Cria um pagamento
// @Description Inicia o processo de pagamento via Stripe Checkout e retorna a URL do Checkout
// @Tags Pagamentos
// @Accept json
// @Produce json
// @Param request body models.Order true "Dados do pedido"
// @Success 200 {object} map[string]string "checkout_url"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao criar pagamento"
// @Router /payments [post]
func CreatePaymentIntent(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	clientSecret, err := services.CreateStripeCheckout(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pagamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"client_secret": clientSecret})
}
