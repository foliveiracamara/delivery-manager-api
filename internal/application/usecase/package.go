package usecase

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/api/http/dto"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/repository"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
	"github.com/foliveiracamara/delivery-manager-api/internal/service"
)

type PackageUseCase struct {
	repository repository.PackageRepository
	service    *service.PackageService
}

func NewPackage(repository repository.PackageRepository, service *service.PackageService) *PackageUseCase {
	return &PackageUseCase{repository: repository, service: service}
}

func (s PackageUseCase) Create(dto dto.PackageRequest) (id string, err error) {
	pkg, err := s.service.Create(&domain.Package{
		Product:           dto.Product,
		WeightKg:          dto.WeightKg,
		DestinationRegion: domain.DestinationRegion(dto.DestinationRegion),
	})
	if err != nil {
		return "", err
	}
	s.repository.Save(pkg)

	return pkg.ID, nil
}

func (s PackageUseCase) Get(id string) (*domain.Package, error) {
	pkg, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func (s PackageUseCase) GetAll() ([]*domain.Package, error) {
	pkgs, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return pkgs, nil
}

func (s PackageUseCase) QuoteShipping(id string) ([]vo.Shipping, error) {
	pkg, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.service.QuoteAvailableShippings(pkg)
}

func (s PackageUseCase) HireCarrier(id string, carrierID string) error {
	pkg, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	s.service.HireCarrier(pkg, carrierID)

	err = s.repository.Save(pkg)
	if err != nil {
		return err
	}

	return nil
}
