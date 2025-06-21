package domain

// CarrierRegion represents the coverage of a carrier in a region
type CarrierRegion struct {
	Region        DestinationRegion `json:"regiao"`
	EstimatedDays int               `json:"prazo_estimado_dias"`
	PricePerKg    float64           `json:"preco_por_kg"`
}

// Carrier represents a carrier in the system
type Carrier struct {
	ID      string          `json:"id"`
	Name    string          `json:"nome"`
	Regions []CarrierRegion `json:"regioes"`
}

// NewCarrier creates a new instance of Carrier
func NewCarrier(id, name string, regions []CarrierRegion) *Carrier {
	return &Carrier{
		ID:      id,
		Name:    name,
		Regions: regions,
	}
}

// GetRegionInfo returns the information of a specific region
func (c *Carrier) GetRegionInfo(region DestinationRegion) (*CarrierRegion, bool) {
	for _, r := range c.Regions {
		if r.Region == region {
			return &r, true
		}
	}
	return nil, false
}

// CalculateShipping calculates the cost and delivery time for a region
func (c *Carrier) CalculateShipping(region DestinationRegion, weightKg float64) (float64, int, bool) {
	regionInfo, exists := c.GetRegionInfo(region)
	if !exists {
		return 0, 0, false
	}

	price := regionInfo.PricePerKg * weightKg
	return price, regionInfo.EstimatedDays, true
}

// IsAvailableForRegion checks if the carrier serves a region
func (c *Carrier) IsAvailableForRegion(region DestinationRegion) bool {
	_, exists := c.GetRegionInfo(region)
	return exists
}

// GetAvailableCarriers returns the available carriers in the system
func GetAvailableCarriers() []*Carrier {
	return []*Carrier{
		// Nebulix Logística
		NewCarrier("nebulix", "Nebulix Logística", []CarrierRegion{
			{
				Region:        DestinationRegionSouth,
				EstimatedDays: 4,
				PricePerKg:    5.90,
			},
			{
				Region:        DestinationRegionSoutheast,
				EstimatedDays: 4,
				PricePerKg:    5.90,
			},
		}),

		// RotaFácil Transportes
		NewCarrier("rotafacil", "RotaFácil Transportes", []CarrierRegion{
			{
				Region:        DestinationRegionSouth,
				EstimatedDays: 7,
				PricePerKg:    4.35,
			},
			{
				Region:        DestinationRegionSoutheast,
				EstimatedDays: 7,
				PricePerKg:    4.35,
			},
			{
				Region:        DestinationRegionMidwest,
				EstimatedDays: 9,
				PricePerKg:    6.22,
			},
			{
				Region:        DestinationRegionNortheast,
				EstimatedDays: 13,
				PricePerKg:    8.00,
			},
		}),

		// Moventra Express
		NewCarrier("moventra", "Moventra Express", []CarrierRegion{
			{
				Region:        DestinationRegionMidwest,
				EstimatedDays: 7,
				PricePerKg:    7.30,
			},
			{
				Region:        DestinationRegionNortheast,
				EstimatedDays: 10,
				PricePerKg:    9.50,
			},
		}),
	}
}
