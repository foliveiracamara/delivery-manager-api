package in_memory

import (
	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/repository"
)

type InMemoryPackageRepository struct {
	packages map[string]*domain.Package
}

func NewInMemoryPackageRepository() repository.PackageRepository {
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

func (r *InMemoryPackageRepository) GetAll() ([]*domain.Package, error) {
	pkgs := make([]*domain.Package, 0, len(r.packages))
	for _, pkg := range r.packages {
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

// Legacy functions for backward compatibility
var packages = make(map[string]*domain.Package)

func Save(pkg *domain.Package) error {
	packages[pkg.ID] = pkg
	return nil
}

func GetPackage(id string) (*domain.Package, error) {
	if pkg, ok := packages[id]; ok {
		return pkg, nil
	}
	return nil, apperr.NewNotFoundError("Package not found")
}
