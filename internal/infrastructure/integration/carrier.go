package integration

import (
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
)

// CarrierRegion represents the coverage of a carrier in a region
type CarrierRegion struct {
	Region        string  `json:"regiao"`
	EstimatedDays int     `json:"prazo_estimado_dias"`
	PricePerKg    float64 `json:"preco_por_kg"`
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
func (c *Carrier) GetRegionInfo(region string) (*CarrierRegion, bool) {
	for _, r := range c.Regions {
		if r.Region == region {
			return &r, true
		}
	}
	return nil, false
}

// CalculateShipping calculates the cost and delivery time for a region
func (c *Carrier) CalculateShipping(region string, weightKg float64) (float64, int, bool) {
	regionInfo, exists := c.GetRegionInfo(region)
	if !exists {
		return 0, 0, false
	}

	price := regionInfo.PricePerKg * weightKg

	// Se o preço for menor que o preço por kg da região, usar o preço por kg da região como valor mínimo
	if price < regionInfo.PricePerKg {
		return regionInfo.PricePerKg, regionInfo.EstimatedDays, true
	}

	return price, regionInfo.EstimatedDays, true
}

// IsAvailableForRegion checks if the carrier serves a region
func (c *Carrier) IsAvailableForRegion(region string) bool {
	_, exists := c.GetRegionInfo(region)
	return exists
}

// GetName returns the carrier name
func (c *Carrier) GetName() string {
	return c.Name
}

// GetID returns the carrier ID
func (c *Carrier) GetID() string {
	return c.ID
}

// CarrierRepository defines the interface for carrier operations
type CarrierRepository interface {
	GetAll() []*Carrier
	GetByID(id string) (*Carrier, error)
}

// CarrierRepositoryImpl implements the carrier repository interface with in-memory storage
// Should simulate a api call to get the carriers
type CarrierRepositoryImpl struct {
	carriers []*Carrier
}

// NewCarrierRepository creates a new instance of CarrierRepository
func NewCarrierRepository() CarrierRepository {
	return &CarrierRepositoryImpl{
		carriers: getAvailableCarriers(),
	}
}

// GetAll returns all available carriers
func (r *CarrierRepositoryImpl) GetAll() []*Carrier {
	return r.carriers
}

// GetByID returns a carrier by its ID
func (r *CarrierRepositoryImpl) GetByID(id string) (*Carrier, error) {
	for _, carrier := range r.carriers {
		if carrier.ID == id {
			return carrier, nil
		}
	}
	return nil, apperr.NewNotFoundError("Carrier not found")
}

// getAvailableCarriers returns the available carriers in the API call
func getAvailableCarriers() []*Carrier {
	return []*Carrier{
		// Nebulix Logística
		NewCarrier("nebulix", "Nebulix Logística", []CarrierRegion{
			{
				Region:        "sul",
				EstimatedDays: 4,
				PricePerKg:    5.90,
			},
			{
				Region:        "sudeste",
				EstimatedDays: 4,
				PricePerKg:    5.90,
			},
		}),

		// RotaFácil Transportes
		NewCarrier("rotafacil", "RotaFácil Transportes", []CarrierRegion{
			{
				Region:        "sul",
				EstimatedDays: 7,
				PricePerKg:    4.35,
			},
			{
				Region:        "sudeste",
				EstimatedDays: 7,
				PricePerKg:    4.35,
			},
			{
				Region:        "centro-oeste",
				EstimatedDays: 9,
				PricePerKg:    6.22,
			},
			{
				Region:        "nordeste",
				EstimatedDays: 13,
				PricePerKg:    8.00,
			},
		}),

		// Moventra Express
		NewCarrier("moventra", "Moventra Express", []CarrierRegion{
			{
				Region:        "centro-oeste",
				EstimatedDays: 7,
				PricePerKg:    7.30,
			},
			{
				Region:        "nordeste",
				EstimatedDays: 10,
				PricePerKg:    9.50,
			},
		}),
	}
}
