// services/stripe_service.go
package stripe_service

import (
	"atlanta-site/models"
	"fmt"
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

func CreateStripeCheckout(order models.Order) (string, error) {
	// Obtém a chave secreta do Stripe
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeKey == "" {
		return "", fmt.Errorf("STRIPE_SECRET_KEY não configurada")
	}

	// Configura a chave de autenticação do Stripe
	stripe.Key = stripeKey

	// Converte o ID do pedido para string
	orderID := strconv.Itoa(int(order.ID))

	// Convertendo o valor para centavos
	amountInCents := int64(order.TotalPrice * 100)

	// Criação do PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents), // Convertendo para centavos
		Currency: stripe.String("brl"),
		Metadata: map[string]string{
			"order_id": orderID, // Usando o ID do pedido como string
		},
	}

	// Criação do PaymentIntent
	intent, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	// Retorna o clientSecret
	return intent.ClientSecret, nil
}
