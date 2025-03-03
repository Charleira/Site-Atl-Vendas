package models

import "time"

// PaymentRequest representa uma solicitação de pagamento.
type PaymentRequest struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`   // ID do pedido associado a esta solicitação de pagamento
	Amount    float64   `json:"amount"`     // Valor do pagamento
	Method    string    `json:"method"`     // Método de pagamento (cartão, boleto, etc.)
	Status    string    `json:"status"`     // Status do pagamento (pendente, pago, falhado)
	CreatedAt time.Time `json:"created_at"` // Data e hora da criação da solicitação
	UpdatedAt time.Time `json:"updated_at"` // Data e hora da última atualização
}
