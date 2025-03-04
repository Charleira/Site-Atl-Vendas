// models/shipping.go
package models

// ShippingOption representa uma opção de envio disponível
// @Description Representa um método de envio disponível via LOGGI
type ShippingOption struct {
	ID    string  `json:"id" example:"1"`
	Name  string  `json:"name" example:"Entrega Expressa"`
	Price float64 `json:"price" example:"19.99"`
	ETA   string  `json:"eta" example:"1-2 dias úteis"`
}

// ShippingRequest representa uma solicitação de criação de envio
// @Description Dados necessários para criar um envio via LOGGI
type ShippingRequest struct {
	OrderID     uint   `json:"order_id" example:"123"`
	Recipient   string `json:"recipient" example:"João Silva"`
	Address     string `json:"address" example:"Rua das Palmeiras, 500"`
	City        string `json:"city" example:"São Paulo"`
	PostalCode  string `json:"postal_code" example:"01010-000"`
	PhoneNumber string `json:"phone_number" example:"+5511999999999"`
}

// ShippingResponse representa a resposta da criação do envio
// @Description Resposta após criação de um envio
type ShippingResponse struct {
	TrackingID string `json:"tracking_id" example:"LOGGI-12345"`
	Status     string `json:"status" example:"Pedido Criado"`
	Estimated  string `json:"estimated" example:"3-5 dias úteis"`
}
