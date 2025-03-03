// services/order_service/order_service.go
package order_service

import (
	"atlanta-site/models"
	"atlanta-site/repositories"
	"errors"
)

// CreateOrderService cria um novo pedido
func CreateOrderService(order *models.Order) error {
	// Chama o repositório para criar o pedido
	if err := repositories.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

// ProcessPaymentService processa o pagamento
func ProcessPaymentService(paymentRequest models.PaymentRequest) error {
	// Aqui você pode implementar a lógica para processar o pagamento
	// Por exemplo, consultar o gateway de pagamento, confirmar o pagamento, etc.
	// Por enquanto, apenas retornamos nil, simulando sucesso no pagamento.

	// Verifique se o valor do pagamento está correto
	if paymentRequest.Amount <= 0 {
		return errors.New("valor do pagamento inválido")
	}

	// Processamento do pagamento (aqui você pode integrar com APIs de pagamento reais)
	// Neste exemplo, assumimos que o pagamento foi bem-sucedido

	return nil
}

// ChangeOrderStatusService altera o status de um pedido
func ChangeOrderStatusService(orderID uint, status string) error {
	// Alterar o status do pedido
	if err := repositories.UpdateOrderStatus(orderID, status); err != nil {
		return err
	}
	return nil
}

// GetOrdersService retorna todos os pedidos
func GetOrdersService() ([]models.Order, error) {
	orders, err := repositories.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
