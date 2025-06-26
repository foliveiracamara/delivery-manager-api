package domain

import (
	"slices"
	"time"

	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
	"github.com/google/uuid"
)

type Package struct {
	ID                string            `json:"id"`
	Product           string            `json:"produto"`
	WeightKg          float64           `json:"peso_kg"`
	DestinationRegion DestinationRegion `json:"regiao_destino"`
	DestinationState  string            `json:"estado_destino"`
	Status            PackageStatus     `json:"status"`
	Shipping          *vo.Shipping      `json:"shipping"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

func NewPackage(product, destinationState string, weightKg float64, destinationRegion DestinationRegion) (*Package, error) {
	valid := isValidDestinationRegion(destinationRegion)
	if !valid {
		return nil, apperr.NewBadRequestError("Invalid destination region")
	}

	now := time.Now()
	pkg := Package{
		ID:                uuid.New().String(),
		Product:           product,
		WeightKg:          weightKg,
		DestinationRegion: destinationRegion,
		DestinationState:  destinationState,
		Status:            StatusCreated,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	return &pkg, nil
}

// UpdateStatus updates the package status
func (p *Package) UpdateStatus(status PackageStatus) error {
	if !IsValidStatus(status) {
		return apperr.NewBadRequestError("Invalid status")
	}

	p.Status = status
	p.UpdatedAt = time.Now()

	return nil
}

func (p *Package) QuoteAvailableShippings() ([]vo.Shipping, error) {
	// Saber qual o destino do pacote
	// Saber quais transportadoras atendem a região
	// Criar lista de cotações de frete
	// Retornar lista de cotações de frete, ordenadas por entrega mais rápida.

	availableCarriers := []*Carrier{}
	for _, carrier := range GetAvailableCarriers() {
		if carrier.IsAvailableForRegion(p.DestinationRegion) {
			availableCarriers = append(availableCarriers, carrier)
		}
	}

	shippings := []vo.Shipping{}
	for _, carrier := range availableCarriers {
		price, days, ok := carrier.CalculateShipping(p.DestinationRegion, p.WeightKg)
		if !ok {
			return nil, apperr.NewInternalServerError("Failed to calculate shipping")
		}
		shippings = append(shippings, vo.NewShippingQuote(
			carrier.Name,
			carrier.ID,
			price,
			days,
		))
	}

	return shippings, nil
}

func (p Package) SortShippingsByDeliveryTime(shippings []vo.Shipping) []vo.Shipping {
	slices.SortFunc(shippings, func(a, b vo.Shipping) int {
		return a.EstimatedDays - b.EstimatedDays
	})
	return shippings
}

func (p *Package) HireCarrier(carrierID string) error {
	// Verificar se o pacote já tem transportadora
	// verificar se a transportadora atende a regiao
	// Se não tiver transportadora, criar uma nova

	if p.Shipping != nil {
		return apperr.NewConflictError("Package already has a carrier")
	}

	carrier, err := GetCarrierByID(carrierID)
	if err != nil {
		return err
	}

	if !carrier.IsAvailableForRegion(p.DestinationRegion) {
		return apperr.NewBadRequestError("Carrier does not serve the destination region")
	}

	price, days, ok := carrier.CalculateShipping(p.DestinationRegion, p.WeightKg)
	if !ok {
		return apperr.NewInternalServerError("Failed to calculate shipping")
	}

	shipping := vo.NewShippingQuote(
		carrier.Name,
		carrier.ID,
		price,
		days,
	)
	p.AssignShipping(shipping)

	return nil
}

// AssignShipping assigns a shipping to the package
func (p *Package) AssignShipping(shipping vo.Shipping) {
	p.Shipping = &shipping
	p.Status = StatusWaitingPickup
	p.UpdatedAt = time.Now()
}

// IsValidStatus checks if the status is valid
func IsValidStatus(status PackageStatus) bool {
	validStatuses := []PackageStatus{
		StatusCreated,
		StatusWaitingPickup,
		StatusCollected,
		StatusShipped,
		StatusDelivered,
		StatusLost,
	}

	return slices.Contains(validStatuses, status)
}

func isValidDestinationRegion(destinationRegion DestinationRegion) bool {
	validRegions := []DestinationRegion{
		DestinationRegionMidwest,
		DestinationRegionNortheast,
		DestinationRegionNorth,
		DestinationRegionSoutheast,
		DestinationRegionSouth,
	}

	return slices.Contains(validRegions, destinationRegion)
}
