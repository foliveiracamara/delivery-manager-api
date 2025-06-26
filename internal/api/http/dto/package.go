package dto

// Requests
type PackageRequest struct {
	Product       string  `json:"produto"`
	WeightKg      float64 `json:"peso_kg"`
	EstadoDestino string  `json:"estado_destino"`
}

type ShippingsQuoteRequest struct {
	PackageID string `json:"package_id"`
}

type HireCarrierRequest struct {
	PackageID string `json:"package_id"`
	CarrierID string `json:"carrier_id"`
}

type UpdateStatusRequest struct {
	PackageID string `json:"package_id"`
	Status    string `json:"status"`
}

// End Requests

// Responses
type PackageResponse struct {
	ID                string                 `json:"id"`
	Product           string                 `json:"produto"`
	WeightKg          float64                `json:"peso_kg"`
	EstadoDestino     string                 `json:"estado_destino"`
	RegiaoDestino     string                 `json:"regiao_destino"`
	Status            string                 `json:"status"`
	Shipping          *ShippingQuoteResponse `json:"entrega,omitempty"`
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
