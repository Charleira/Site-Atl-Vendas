package models

// Cart representa o carrinho de um usuário
type Cart struct {
	UserID   int           `json:"user_id"`
	Products []CartProduct `json:"products"`
}

// CartProduct representa um produto no carrinho
type CartProduct struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
