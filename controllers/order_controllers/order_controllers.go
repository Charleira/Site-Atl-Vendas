// order_controllers/order_controller.go
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

	// Chama o repositório (ou serviço) para criar o pedido
	if err := services.CreateOrderService(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pedido"})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// ProcessPayment processa o pagamento de um pedido
func ProcessPayment(c *gin.Context) {
	var paymentRequest models.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o serviço para processar o pagamento
	if err := services.ProcessPaymentService(paymentRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar pagamento"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pagamento processado com sucesso"})
}

// Define a struct para o body da requisição no UpdateOrderStatus
type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

// UpdateOrderStatus atualiza o status do pedido
func UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do pedido inválido"})
		return
	}

	var request UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Chama o serviço para alterar o status
	if err := services.ChangeOrderStatusService(uint(id), request.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Status atualizado com sucesso"})
}

// ListOrders retorna todos os pedidos
func ListOrders(c *gin.Context) {
	orders, err := services.GetOrdersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
