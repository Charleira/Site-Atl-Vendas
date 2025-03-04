package order_controllers

import (
	"net/http"
	"strconv"

	"atlanta-site/models"
	services "atlanta-site/services/order_service"

	"github.com/gin-gonic/gin"
)

// CreateOrder cria um novo pedido
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
	c.JSON(http.StatusCreated, order)
}

// GetOrderDetails retorna detalhes de um pedido específico
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
