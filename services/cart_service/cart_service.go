package cart_service

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"database/sql"
	"log"
)

// Adiciona um produto ao carrinho do usuário
func AddProductToCart(db *sql.DB, userID int, productID int, quantity int) error {
	err := repositories.AddProductToCart(db, userID, productID, quantity)
	if err != nil {
		log.Println("Erro ao adicionar produto ao carrinho:", err)
		return err
	}
	return nil
}

// Recupera o carrinho de um usuário
func GetCartByUserID(db *sql.DB, userID int) ([]models.CartProduct, error) {
	cart, err := repositories.GetCartByUserID(db, userID)
	if err != nil {
		log.Println("Erro ao recuperar carrinho:", err)
		return nil, err
	}
	return cart, nil
}

// Remove um produto do carrinho
func RemoveProductFromCart(db *sql.DB, userID int, productID int) error {
	err := repositories.RemoveProductFromCart(db, userID, productID)
	if err != nil {
		log.Println("Erro ao remover produto do carrinho:", err)
		return err
	}
	return nil
}
