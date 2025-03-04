// services/stripe_service.go
package stripe_service

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"os"

	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func CreateStripeCheckout(order models.Order) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	checkoutSession, err := session.New(&stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("brl"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(order.Item),
					},
					UnitAmount: stripe.Int64(int64(order.TotalPrice * 100)),
				},
				Quantity: stripe.Int64(int64(order.Quantity)),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://seusite.com/sucesso"),
		CancelURL:  stripe.String("https://seusite.com/cancelado"),
	})

	if err != nil {
		return "", err
	}

	return checkoutSession.URL, nil
}

func UpdatePaymentStatus(stripeID string, status string) error {
	return repositories.UpdatePaymentStatus(stripeID, status)

}
