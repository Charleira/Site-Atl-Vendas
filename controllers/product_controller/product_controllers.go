// controllers/product_controllers.go
package product_controllers

import (
	"atlanta-site/models"
	services "atlanta-site/services/product_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListProducts retorna todos os produtos disponíveis
// @Summary Lista todos os produtos
// @Description Retorna uma lista de produtos disponíveis no catálogo
// @Tags Produtos
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string "Erro ao listar produtos"
// @Router /products [get]
func ListProducts(c *gin.Context) {
	products, err := services.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar produtos"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductById retorna um produto específico pelo ID
// @Summary Obtém um produto
// @Description Retorna um produto pelo seu ID
// @Tags Produtos
// @Accept json
// @Produce json
// @Param product_id path int true "ID do produto"
// @Success 200 {object} models.Product
// @Failure 500 {object} map[string]string "Erro ao recuperar produto"
// @Router /products/{product_id} [get]
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
// @Summary Cria um produto
// @Description Cria um novo produto enviando dados e uma imagem
// @Tags Produtos
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nome do produto"
// @Param description formData string true "Descrição do produto"
// @Param price formData number true "Preço do produto"
// @Param image formData file true "Imagem do produto"
// @Success 201 {object} map[string]string "Produto criado com sucesso"
// @Failure 400 {object} map[string]string "Erro ao receber imagem"
// @Failure 500 {object} map[string]string "Erro ao criar produto"
// @Router /products [post]
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
// @Summary Atualiza um produto
// @Description Atualiza as informações de um produto pelo ID
// @Tags Produtos
// @Accept multipart/form-data
// @Produce json
// @Param product_id path int true "ID do produto"
// @Param name formData string false "Nome do produto"
// @Param description formData string false "Descrição do produto"
// @Param price formData number false "Preço do produto"
// @Param image formData file false "Nova imagem do produto"
// @Success 200 {object} map[string]string "Produto atualizado com sucesso"
// @Failure 500 {object} map[string]string "Erro ao atualizar produto"
// @Router /products/{product_id} [put]
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
// @Summary Remove um produto
// @Description Deleta um produto pelo ID
// @Tags Produtos
// @Accept json
// @Produce json
// @Param product_id path int true "ID do produto"
// @Success 200 {object} map[string]string "Produto removido com sucesso"
// @Failure 500 {object} map[string]string "Erro ao remover produto"
// @Router /products/{product_id} [delete]
func RemoveProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("product_id"))

	err := services.RemoveProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover produto"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produto removido com sucesso"})
}
