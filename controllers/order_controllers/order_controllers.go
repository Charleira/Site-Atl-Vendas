// order_controllers/order_controllers.go
package order_controllers

import (
	"net/http"
	"strconv"
	"bytes"
	"atlanta-site/models"
	services "atlanta-site/services/order_service"

	"github.com/gin-gonic/gin"
)

// CreateOrder cria um novo pedido
// @Summary Cria um novo pedido
// @Description Cria um novo pedido com os dados fornecidos
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body models.Order true "Dados do pedido"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao criar pedido"
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := services.CreateOrderService(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}
	SendDiscordNotification(user.Username, product.Name)
	c.JSON(http.StatusCreated, order)
}

// GetOrderDetails retorna detalhes de um pedido específico
// @Summary Obtém detalhes de um pedido
// @Description Retorna os detalhes de um pedido específico pelo ID
// @Tags Orders
// @Produce json
// @Param order_id path int true "ID do Pedido"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 500 {object} map[string]string "Erro ao buscar detalhes do pedido"
// @Router /orders/{order_id} [get]
func GetOrderDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	order, err := services.GetOrderDetailsService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar detalhes do pedido"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// CancelOrder cancela um pedido
// @Summary Cancela um pedido
// @Description Cancela um pedido específico pelo ID
// @Tags Orders
// @Produce json
// @Param order_id path int true "ID do Pedido"
// @Success 200 {object} map[string]string "Pedido cancelado com sucesso"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 500 {object} map[string]string "Erro ao cancelar pedido"
// @Router /orders/{order_id}/cancel [post]
func CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := services.CancelOrderService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cancelar pedido"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pedido cancelado com sucesso"})
}

// ListOrders lista pedidos do usuário
// @Summary Lista pedidos do usuário
// @Description Retorna uma lista de pedidos de um usuário específico pelo ID
// @Tags Orders
// @Produce json
// @Param user_id path int true "ID do Usuário"
// @Success 200 {array} models.Order
// @Failure 400 {object} map[string]string "ID do usuário inválido"
// @Failure 500 {object} map[string]string "Erro ao buscar pedidos"
// @Router /users/{user_id}/orders [get]
func ListOrders(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do usuário inválido"})
		return
	}

	orders, err := services.ListOrdersService(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// TrackOrder rastreia o status do pedido
// @Summary Rastreia um pedido
// @Description Obtém o status atual de um pedido específico pelo ID
// @Tags Orders
// @Produce json
// @Param order_id path int true "ID do Pedido"
// @Success 200 {object} map[string]string "status"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 500 {object} map[string]string "Erro ao rastrear pedido"
// @Router /orders/{order_id}/track [get]
func TrackOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	status, err := services.TrackOrderService(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao rastrear pedido"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func SendDiscordNotification(username, product string) {
	webhookURL := "https://discord.com/api/webhooks/SEU-WEBHOOK"

	message := fmt.Sprintf("🛒 **%s** comprou **%s**!", username, product)

	payload := map[string]string{"content": message}
	jsonData, _ := json.Marshal(payload)

	http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
}
