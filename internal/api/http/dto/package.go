package dto

// Requests
type PackageRequest struct {
	Product       string  `json:"produto" validate:"required,min=2,max=100"`
	WeightKg      float64 `json:"peso_kg" validate:"required,gt=0,lte=1000"`
	EstadoDestino string  `json:"estado_destino" validate:"required,len=2,alpha"`
}

type ShippingsQuoteRequest struct {
	PackageID string `json:"package_id" validate:"required"`
}

type HireCarrierRequest struct {
	PackageID string `json:"package_id" validate:"required"`
	CarrierID string `json:"carrier_id" validate:"required"`
}

type UpdateStatusRequest struct {
	PackageID string `json:"package_id" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

// End Requests

// Responses
type PackageResponse struct {
	ID            string                 `json:"id"`
	Product       string                 `json:"produto"`
	WeightKg      float64                `json:"peso_kg"`
	EstadoDestino string                 `json:"estado_destino"`
	RegiaoDestino string                 `json:"regiao_destino"`
	Status        string                 `json:"status"`
	Shipping      *ShippingQuoteResponse `json:"entrega,omitempty"`
}

type ShippingsQuoteResponse struct {
	Shippings []ShippingQuoteResponse `json:"transportadoras"`
}

type ShippingQuoteResponse struct {
	CarrierID     string  `json:"transportadora_id"`
	EstimatedDays int     `json:"prazo_estimado"`
	Price         float64 `json:"preco"`
	CarrierName   string  `json:"nome"`
}

// End Responses
