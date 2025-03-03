package repositories

import (
	"atlanta-site/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Adiciona um produto ao carrinho do usuário
func AddProductToCart(db *sql.DB, userID int, productID int, quantity int) error {
	query := `INSERT INTO carts (user_id, product_id, quantity)
              VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE quantity = quantity + ?`
	_, err := db.Exec(query, userID, productID, quantity, quantity)
	return err
}

// Recupera o carrinho de um usuário
func GetCartByUserID(db *sql.DB, userID int) ([]models.CartProduct, error) {
	query := `SELECT product_id, quantity FROM carts WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart []models.CartProduct
	for rows.Next() {
		var cartProduct models.CartProduct
		if err := rows.Scan(&cartProduct.ProductID, &cartProduct.Quantity); err != nil {
			return nil, err
		}
		cart = append(cart, cartProduct)
	}
	return cart, nil
}

// Remove um produto do carrinho
func RemoveProductFromCart(db *sql.DB, userID int, productID int) error {
	query := `DELETE FROM carts WHERE user_id = ? AND product_id = ?`
	_, err := db.Exec(query, userID, productID)
	return err
}
