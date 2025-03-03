package repositories

import (
	"atlanta-site/config"
	"atlanta-site/models"
)

// CreateOrder cria um novo pedido no banco de dados
func CreateOrder(order *models.Order) error {
	query := `INSERT INTO orders (user_id, total, status) VALUES (?, ?, ?)`
	_, err := config.DB.Exec(query, order.UserID, order.TotalPrice, order.Status)
	if err != nil {
		return err
	}
	return nil
}

// GetAllOrders retorna todos os pedidos cadastrados
func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, user_id, total, status FROM orders`
	rows, err := config.DB.Query(query)
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// UpdateOrderStatus altera o status de um pedido
func UpdateOrderStatus(id uint, status string) error {
	query := `UPDATE orders SET status = ? WHERE id = ?`
	_, err := config.DB.Exec(query, status, id)
	return err
}
