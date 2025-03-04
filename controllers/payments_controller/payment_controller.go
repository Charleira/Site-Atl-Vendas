// controllers/payment_controller.go
package payment_controller

import (
	"atlanta-site/models"
	services "atlanta-site/services/stripe_service"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/webhook"
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

	checkoutURL, err := services.CreateStripeCheckout(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pagamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkout_url": checkoutURL})
}

// WebhookPaymentStatus recebe atualizações de status de pagamento do Stripe
// @Summary Webhook de pagamento
// @Description Recebe notificações do Stripe sobre o status do pagamento
// @Tags Pagamentos
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Webhook recebido"
// @Failure 400 {object} map[string]string "Falha ao processar webhook"
// @Router /payments/webhook [post]
func WebhookPaymentStatus(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao ler payload"})
		return
	}

	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	sigHeader := c.Request.Header.Get("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Assinatura inválida"})
		return
	}

	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar webhook"})
			return
		}
		services.UpdatePaymentStatus(session.ID, "pago")
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook recebido"})
}
