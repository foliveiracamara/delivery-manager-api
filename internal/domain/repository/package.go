package repository

import "github.com/foliveiracamara/delivery-manager-api/internal/domain"

type PackageRepository interface {
	Save(pkg *domain.Package) error
	GetByID(id string) (*domain.Package, error)	
	GetAll() ([]*domain.Package, error)
}