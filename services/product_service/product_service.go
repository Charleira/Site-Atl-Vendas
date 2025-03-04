package services

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"errors"
	"mime/multipart"
	"os"
)

// ListProducts retorna todos os produtos
func ListProducts() ([]models.Product, error) {
	return repositories.ListProducts()
}

// GetProductByID retorna um produto pelo ID
func GetProductByID(productID int) (models.Product, error) {
	return repositories.GetProductByID(productID)
}

// CreateProduct adiciona um novo produto com imagem
func CreateProduct(product models.Product, file *multipart.FileHeader) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("dados inválidos para criar produto")
	}

	// Salva a imagem no servidor
	imagePath, err := saveImage(file)
	if err != nil {
		return err
	}
	product.ImageURL = imagePath // Salva o caminho da imagem no banco

	return repositories.CreateProduct(product)
}

// UpdateProduct atualiza os dados do produto
func UpdateProduct(productID int, product models.Product, file *multipart.FileHeader) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("dados inválidos para atualizar produto")
	}

	// Se houver uma nova imagem, substituímos a antiga
	if file != nil {
		imagePath, err := saveImage(file)
		if err != nil {
			return err
		}
		product.ImageURL = imagePath
	}

	return repositories.UpdateProduct(productID, product)
}

// RemoveProduct remove um produto pelo ID
func RemoveProduct(productID int) error {
	return repositories.RemoveProduct(productID)
}

// Função auxiliar para salvar a imagem no servidor
func saveImage(file *multipart.FileHeader) (string, error) {
	filePath := "uploads/" + file.Filename

	// Garante que a pasta de uploads existe
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", err
	}

	// Salva o arquivo no servidor
	err := os.Rename(file.Filename, filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
