package dto

// Requests
type PackageRequest struct {
	Product           string  `json:"product"`
	WeightKg          float64 `json:"weight_kg"`
	DestinationRegion string  `json:"destination_region"`
}

type ShippingsQuoteRequest struct {
	PackageID string `json:"package_id"`
}

type HireCarrierRequest struct {
	PackageID  string `json:"package_id"`
	CarrierID  string `json:"carrier_id"`
}

// End Requests

// Responses
type PackageResponse struct {
	ID                string                 `json:"id"`
	Product           string                 `json:"product"`
	WeightKg          float64                `json:"weight_kg"`
	DestinationRegion string                 `json:"destination_region"`
	Status            string                 `json:"status"`
	Shipping          *ShippingQuoteResponse `json:"shipping,omitempty"`
}

type ShippingsQuoteResponse struct {
	Shippings []ShippingQuoteResponse `json:"shippings"`
}

type ShippingQuoteResponse struct {
	CarrierID     string  `json:"carrier_id"`
	EstimatedDays int     `json:"estimated_days"`
	Price         float64 `json:"price"`
	CarrierName   string  `json:"carrier_name"`
}

// End Responses
