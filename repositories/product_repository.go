package repositories

import (
	"atlanta-site/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Recupera um produto pelo ID
func GetProductByID(db *sql.DB, productID int) (models.Product, error) {
	query := `SELECT id, name, description, price FROM products WHERE id = ?`
	row := db.QueryRow(query, productID)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
