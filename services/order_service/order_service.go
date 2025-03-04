package order_service

import (
	"atlanta-site/config"
	"atlanta-site/models"
	"atlanta-site/repositories"
	"errors"
	"time"
)

// CreateOrderService cria um novo pedido
func CreateOrderService(order *models.Order) error {
	// Definindo o status inicial como "pending"
	order.Status = models.StatusPending

	// Preenche o CreatedAt com a data atual
	order.CreatedAt = time.Now()

	// Insere o pedido no banco de dados
	query := `INSERT INTO orders (item, quantity, price, total_price, user_id, nickname, size, status, created_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := config.DB.Exec(query, order.Item, order.Quantity, order.Price, order.TotalPrice, order.UserID, order.Nickname, order.Size, order.Status, order.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderDetailsService busca os detalhes de um pedido
func GetOrderDetailsService(orderID uint) (*models.Order, error) {
	return repositories.GetOrderById(orderID)
}

// CancelOrderService cancela um pedido se permitido
func CancelOrderService(orderID uint) error {
	order, err := repositories.GetOrderById(orderID)
	if err != nil {
		return err
	}

	if order.Status == "Cancelado" {
		return errors.New("pedido já cancelado")
	}

	return repositories.CancelOrder(orderID)
}

// ListOrdersService retorna pedidos do usuário
func ListOrdersService(userID uint) ([]models.Order, error) {
	return repositories.GetOrdersByUser(userID)
}

// TrackOrderService retorna o status de um pedido
func TrackOrderService(orderID uint) (string, error) {
	order, err := repositories.GetOrderById(orderID)
	if err != nil {
		return "", err
	}
	return order.Status, nil
}
