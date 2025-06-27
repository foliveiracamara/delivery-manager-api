package usecase

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/api/http/dto"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/vo"
	"github.com/foliveiracamara/delivery-manager-api/internal/service"
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
)

type PackageUseCase struct {
	repository domain.PackageRepository
	service    *service.PackageService
}

func NewPackage(repository domain.PackageRepository, service *service.PackageService) *PackageUseCase {
	return &PackageUseCase{
		repository: repository,
		service:    service,
	}
}

func (s PackageUseCase) Create(dto dto.PackageRequest) (id string, err error) {
	// Convert state to region
	region, exists := domain.GetRegionFromState(dto.EstadoDestino)
	if !exists {
		return "", apperr.NewBadRequestError("Invalid state: " + dto.EstadoDestino)
	}

	pkg, err := s.service.Create(&domain.Package{
		Product:           dto.Product,
		WeightKg:          dto.WeightKg,
		DestinationRegion: region,
		DestinationState:  dto.EstadoDestino,
	})
	if err != nil {
		return "", err
	}

	err = s.repository.Save(pkg)
	if err != nil {
		return "", err
	}

	return pkg.ID, nil
}

func (s PackageUseCase) Get(id string) (*domain.Package, error) {
	pkg, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func (s PackageUseCase) UpdateStatus(id string, status string) error {
	pkg, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	err = s.service.UpdateStatus(pkg, domain.PackageStatus(status))
	if err != nil {
		return err
	}

	return s.repository.Save(pkg)
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

	err = s.service.HireCarrier(pkg, carrierID)
	if err != nil {
		return err
	}

	err = s.repository.Save(pkg)
	if err != nil {
		return err
	}

	return nil
}
