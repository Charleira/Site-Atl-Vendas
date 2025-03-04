// models/shipping.go
package models

// ShippingOption representa uma opção de envio disponível
type ShippingOption struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	ETA   string  `json:"eta"` // Estimativa de tempo de entrega
}

// ShippingRequest representa uma solicitação de criação de envio
type ShippingRequest struct {
	OrderID     uint   `json:"order_id"`
	Recipient   string `json:"recipient"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	PhoneNumber string `json:"phone_number"`
}

// ShippingResponse representa a resposta da criação do envio
type ShippingResponse struct {
	TrackingID string `json:"tracking_id"`
	Status     string `json:"status"`
	Estimated  string `json:"estimated"`
}
