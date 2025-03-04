package models

// Product representa um produto disponível no sistema
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"` // Novo campo para armazenar o caminho da imagem
}
