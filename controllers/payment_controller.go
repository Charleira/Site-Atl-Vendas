package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"atlanta-site/models"
)

// ProcessPayment inicia o Stripe Checkout e retorna a URL para pagamento
func ProcessPayment(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Inicializa a API do Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Cria uma sessão do Checkout
	checkoutSession, err := session.New(&stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("brl"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(order.Item),
					},
					UnitAmount: stripe.Int64(int64(order.TotalPrice * 100)), // Stripe usa centavos
				},
				Quantity: stripe.Int64(int64(order.Quantity)),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://seusite.com/sucesso"),
		CancelURL:  stripe.String("https://seusite.com/cancelado"),
	})

	if err != nil {
		log.Println("Erro ao criar sessão do Stripe:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar pagamento"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkout_url": checkoutSession.URL})
}
