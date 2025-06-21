package application

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/api/http"
	"github.com/foliveiracamara/delivery-manager-api/internal/application/usecase"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain/repository"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/config"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/persistence"
	"github.com/foliveiracamara/delivery-manager-api/internal/service"
	"go.uber.org/fx"
)

func providers() []any {
	globalDeps := []any{
		config.New,

		// Repository
		fx.Annotate(persistence.NewInMemoryPackageRepository, fx.As(new(repository.PackageRepository))),

		// Services
		service.NewPackageService,

		// Use Cases
		usecase.NewPackage,

		// HTTP
		http.NewControllerManager,
	}

	globalDeps = append(globalDeps, http.ControllersList...)
	return globalDeps
}
