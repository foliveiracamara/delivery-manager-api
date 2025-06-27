package domain

type PackageRepository interface {
	Save(pkg *Package) error
	GetByID(id string) (*Package, error)
}
