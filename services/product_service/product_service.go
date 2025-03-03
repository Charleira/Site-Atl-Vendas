package product_service

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"database/sql"
	"log"
)

// Recupera um produto pelo ID
func GetProductByID(db *sql.DB, productID int) (models.Product, error) {
	product, err := repositories.GetProductByID(db, productID)
	if err != nil {
		log.Println("Erro ao recuperar produto:", err)
		return models.Product{}, err
	}
	return product, nil
}
