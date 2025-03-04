package repositories

import (
	"atlanta-site/config"
	"atlanta-site/models"
	"database/sql"
)

// ListProducts retorna todos os produtos disponíveis
func ListProducts() ([]models.Product, error) {
	query := `SELECT id, name, description, price, image_url FROM products`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// GetProductByID busca um produto pelo ID
func GetProductByID(productID int) (models.Product, error) {
	query := `SELECT id, name, description, price, image_url FROM products WHERE id = ?`
	row := config.DB.QueryRow(query, productID)

	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL)
	if err == sql.ErrNoRows {
		return models.Product{}, nil
	}
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// CreateProduct insere um novo produto no banco
func CreateProduct(product models.Product) error {
	query := `INSERT INTO products (name, description, price, image_url) VALUES (?, ?, ?, ?)`
	_, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.ImageURL)
	return err
}

// UpdateProduct atualiza um produto existente
func UpdateProduct(productID int, product models.Product) error {
	query := `UPDATE products SET name = ?, description = ?, price = ?, image_url = ? WHERE id = ?`
	_, err := config.DB.Exec(query, product.Name, product.Description, product.Price, product.ImageURL, productID)
	return err
}

// RemoveProduct remove um produto pelo ID
func RemoveProduct(productID int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := config.DB.Exec(query, productID)
	return err
}
