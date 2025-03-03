package cart_controllers

import (
	"atlanta-site/config"
	services "atlanta-site/services/cart_service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Adiciona um produto ao carrinho
// @Summary Adiciona um produto ao carrinho
// @Description Adiciona um produto ao carrinho do usuário
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "ID do Usuário"
// @Param product_id path int true "ID do Produto"
// @Param quantity query int false "Quantidade do Produto"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/{user_id}/add/{product_id}/ [post]
func AddProductToCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))
	quantity, _ := strconv.Atoi(c.DefaultQuery("quantity", "1"))

	err := services.AddProductToCart(config.DB, userID, productID, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar produto ao carrinho"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produto adicionado ao carrinho"})
}

// Recupera o carrinho do usuário
// @Summary Recupera o carrinho do usuário
// @Description Retorna os produtos do carrinho do usuário
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "ID do Usuário"
// @Success 200 {array} models.CartProduct
// @Failure 500 {object} map[string]string
// @Router /cart/{user_id}/ [get]
func GetCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	cart, err := services.GetCartByUserID(&sql.DB{}, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar carrinho"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// Remove um produto do carrinho
// @Summary Remove um produto do carrinho
// @Description Remove um produto do carrinho do usuário
// @Tags Cart
// @Accept json
// @Produce json
// @Param user_id path int true "ID do Usuário"
// @Param product_id path int true "ID do Produto"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cart/{user_id}/remove/{product_id}/ [delete]
func RemoveProductFromCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))

	err := services.RemoveProductFromCart(&sql.DB{}, userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover produto do carrinho"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produto removido do carrinho"})
}
