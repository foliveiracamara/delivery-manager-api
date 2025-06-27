package vo

// Shipping representa uma cotação de frete
type Shipping struct {
	CarrierName    string
	EstimatedPrice float64
	EstimatedDays  int
	CarrierID      string
}

// ShippingRequest representa uma requisição de cotação
type ShippingRequest struct {
	WeightKg         float64
	DestinationState string
}

// NewShippingQuote cria uma nova cotação de frete
func NewShippingQuote(carrierName, carrierID string, estimatedPrice float64, estimatedDays int) Shipping {
	return Shipping{
		CarrierName:    carrierName,
		EstimatedPrice: estimatedPrice,
		EstimatedDays:  estimatedDays,
		CarrierID:      carrierID,
	}
}

// NewShippingRequest cria uma nova requisição de cotação
func NewShippingRequest(weightKg float64, destinationState string) ShippingRequest {
	return ShippingRequest{
		WeightKg:         weightKg,
		DestinationState: destinationState,
	}
}

// IsValid verifica se a requisição de cotação é válida
func (sr ShippingRequest) IsValid() bool {
	return sr.WeightKg > 0 && sr.DestinationState != ""
}