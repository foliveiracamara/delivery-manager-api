package persistence

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
)

type InMemoryPackageRepository struct {
	packages map[string]*domain.Package
}

func NewInMemoryPackageRepository() domain.PackageRepository {
	return &InMemoryPackageRepository{
		packages: make(map[string]*domain.Package),
	}
}

func (r *InMemoryPackageRepository) Save(pkg *domain.Package) error {
	r.packages[pkg.ID] = pkg
	return nil
}

func (r *InMemoryPackageRepository) GetByID(id string) (*domain.Package, error) {
	if pkg, ok := r.packages[id]; ok {
		return pkg, nil
	}
	return nil, apperr.NewNotFoundError("Package not found")
}
