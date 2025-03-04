// services/shipping_service.go
package shipping_service

import (
	"atlanta-site/models"
	"errors"
	"fmt"
)

// Simulação de integração com a API da LOGGI
func GetShippingOptions() ([]models.ShippingOption, error) {
	// Aqui deveria chamar a API da LOGGI para obter opções reais
	options := []models.ShippingOption{
		{ID: "1", Name: "Entrega Rápida", Price: 19.99, ETA: "1-2 dias úteis"},
		{ID: "2", Name: "Entrega Padrão", Price: 9.99, ETA: "3-5 dias úteis"},
	}
	return options, nil
}

func CreateShipping(request models.ShippingRequest) (*models.ShippingResponse, error) {
	// Simulação de chamada para a API da LOGGI para criar um envio
	if request.OrderID == 0 || request.Recipient == "" || request.Address == "" {
		return nil, errors.New("dados do envio incompletos")
	}

	// Aqui faria a chamada real à API da LOGGI para criar o pedido de entrega
	trackingID := fmt.Sprintf("LOGGI-%d", request.OrderID)

	response := &models.ShippingResponse{
		TrackingID: trackingID,
		Status:     "Pedido Criado",
		Estimated:  "3-5 dias úteis",
	}

	return response, nil
}
