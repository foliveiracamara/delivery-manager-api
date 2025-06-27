package service

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/integration"
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
)

// PackageService represents the package service
type PackageService struct {
	carrierRepo integration.CarrierRepository
}

// NewPackageService creates a new instance of PackageService
func NewPackageService(carrierRepo integration.CarrierRepository) *PackageService {
	return &PackageService{
		carrierRepo: carrierRepo,
	}
}

func (s PackageService) Create(pkg *domain.Package) (*domain.Package, error) {
	pkg, err := domain.NewPackage(
		pkg.Product,
		pkg.DestinationState,
		pkg.WeightKg,
		pkg.DestinationRegion,
	)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

func (s PackageService) UpdateStatus(pkg *domain.Package, status domain.PackageStatus) error {
	return pkg.UpdateStatus(status)
}

func (s PackageService) QuoteAvailableShippings(pkg *domain.Package) ([]vo.Shipping, error) {
	availableCarriers := []*integration.Carrier{}
	allCarriers := s.carrierRepo.GetAll()
	destinationRegion := string(pkg.DestinationRegion)

	for _, carrier := range allCarriers {
		if carrier.IsAvailableForRegion(destinationRegion) {
			availableCarriers = append(availableCarriers, carrier)
		}
	}

	shippings := []vo.Shipping{}
	for _, carrier := range availableCarriers {
		price, days, ok := carrier.CalculateShipping(destinationRegion, pkg.WeightKg)
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

	sortedShippings := pkg.SortShippingsByDeliveryTime(shippings)
	return sortedShippings, nil
}

func (s PackageService) HireCarrier(pkg *domain.Package, carrierID string) error {
	if pkg.Shipping != nil {
		return apperr.NewConflictError("Package already has a carrier")
	}

	carrier, err := s.carrierRepo.GetByID(carrierID)
	if err != nil {
		return err
	}

	destinationRegion := string(pkg.DestinationRegion)
	if !carrier.IsAvailableForRegion(destinationRegion) {
		return apperr.NewBadRequestError("Carrier does not serve the destination region")
	}

	price, days, ok := carrier.CalculateShipping(destinationRegion, pkg.WeightKg)
	if !ok {
		return apperr.NewInternalServerError("Failed to calculate shipping")
	}

	shipping := vo.NewShippingQuote(
		carrier.Name,
		carrier.ID,
		price,
		days,
	)
	pkg.AssignShipping(shipping)

	return nil
}
