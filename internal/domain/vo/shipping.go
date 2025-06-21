package vo

// Shipping representa uma cotação de frete (Value Object)
type Shipping struct {
	CarrierName    string  `json:"transportadora"`
	EstimatedPrice float64 `json:"preco_estimado"`
	EstimatedDays  int     `json:"prazo_estimado_dias"`
	CarrierID      string  `json:"transportadora_id"`
}

// ShippingRequest representa uma requisição de cotação (Value Object)
type ShippingRequest struct {
	WeightKg         float64 `json:"peso_kg"`
	DestinationState string  `json:"estado_destino"`
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

// TODO: Checar se vai precisar mesmo
// GetRegion retorna a região baseada no estado de destino
// func (sr ShippingRequest) GetRegion() Region {
// 	return GetRegionFromState(sr.DestinationState)
// }
