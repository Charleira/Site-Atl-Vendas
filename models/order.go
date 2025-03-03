package models

import (
	"time"
)

// Definição dos possíveis status do pedido
const (
	StatusPending   = "pending"
	StatusPaid      = "paid"
	StatusShipped   = "shipped"
	StatusDelivered = "delivered"
	StatusCanceled  = "canceled"
)

// Order representa um pedido no sistema
type Order struct {
	ID         uint      `json:"id"`
	Item       string    `json:"item"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
	UserID     uint      `json:"user_id"`
	Nickname   string    `json:"nickname"`
	Size       string    `json:"size"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
	DeletedAt  time.Time `json:"-"` // Soft delete, usando time.Time para o campo DeletedAt
}
