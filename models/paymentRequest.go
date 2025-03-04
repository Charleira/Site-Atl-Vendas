// models/payment.go
package models

import "time"

type Payment struct {
	ID        uint      `json:"id"`
	OrderID   uint      `json:"order_id"`
	StripeID  string    `json:"stripe_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
