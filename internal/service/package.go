package service

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
)

// TODO: Criar interface para o service
type PackageService struct{}

func NewPackageService() *PackageService {
	return &PackageService{}
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
	shippings, err := pkg.QuoteAvailableShippings()
	if err != nil {
		return nil, err
	}
	sortedShippings := pkg.SortShippingsByDeliveryTime(shippings)

	return sortedShippings, nil
}

func (s PackageService) HireCarrier(pkg *domain.Package, carrierID string) error {
	return pkg.HireCarrier(carrierID)
}