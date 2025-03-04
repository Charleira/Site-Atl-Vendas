// repositories/payment_repository.go
package repositories

import (
	"atlanta-site/config"
	"atlanta-site/models"
	"database/sql"
	"time"
)

// Criar pagamento no banco
func CreatePayment(orderID uint, stripeID string, amount float64, status string) error {
	query := `INSERT INTO payments (order_id, stripe_id, amount, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, orderID, stripeID, amount, status, time.Now(), time.Now())
	return err
}

// Atualizar status do pagamento
func UpdatePaymentStatus(stripeID string, status string) error {
	query := `UPDATE payments SET status = ?, updated_at = ? WHERE stripe_id = ?`
	_, err := config.DB.Exec(query, status, time.Now(), stripeID)
	return err
}

// Obter pagamento por Stripe ID
func GetPaymentByStripeID(stripeID string) (*models.Payment, error) {
	var payment models.Payment
	query := `SELECT id, order_id, stripe_id, amount, status, created_at, updated_at FROM payments WHERE stripe_id = ?`
	row := config.DB.QueryRow(query, stripeID)
	err := row.Scan(&payment.ID, &payment.OrderID, &payment.StripeID, &payment.Amount, &payment.Status, &payment.CreatedAt, &payment.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil
}
