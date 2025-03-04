package product_controllers

import (
	"atlanta-site/models"
	services "atlanta-site/services/product_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListProducts retorna todos os produtos disponíveis
func ListProducts(c *gin.Context) {
	products, err := services.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar produtos"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductById retorna um produto específico pelo ID
func GetProductById(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("product_id"))
	product, err := services.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar produto"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct cria um novo produto com imagem
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Obtém os dados do produto do formulário
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	product.Price = price

	// Obtém o arquivo de imagem enviado
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao receber imagem"})
		return
	}

	// Salva o produto no banco de dados
	err = services.CreateProduct(product, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produto criado com sucesso"})
}

// UpdateProduct atualiza os dados de um produto
func UpdateProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	var product models.Product
	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	product.Price = price

	file, _ := c.FormFile("image") // A imagem pode ser opcional

	err := services.UpdateProduct(productID, product, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto atualizado com sucesso"})
}

// RemoveProduct exclui um produto
func RemoveProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	err := services.RemoveProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
