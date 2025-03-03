package product_controllers

import (
	services "atlanta-site/services/product_service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Recupera um produto pelo ID
// @Summary Recupera um produto pelo ID
// @Description Retorna os detalhes de um produto específico
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path int true "ID do Produto"
// @Success 200 {object} models.Product
// @Failure 500 {object} map[string]string
// @Router /product/{product_id}/ [get]
func GetProduct(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("product_id"))
	product, err := services.GetProductByID(&sql.DB{}, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar produto"})
		return
	}
	c.JSON(http.StatusOK, product)
}
