package repositories

import (
	"atlanta-site/config"
	"atlanta-site/models"
)

// GetOrderById retorna um pedido específico
func GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	query := `SELECT id, user_id, total, status FROM orders WHERE id = ?`
	err := config.DB.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// CancelOrder atualiza o status do pedido para "Cancelado"
func CancelOrder(id uint) error {
	query := `UPDATE orders SET status = 'Cancelado' WHERE id = ?`
	_, err := config.DB.Exec(query, id)
	return err
}

// GetOrdersByUser retorna pedidos de um usuário
func GetOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, user_id, total, status FROM orders WHERE user_id = ?`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
