package application

import (
	"github.com/foliveiracamara/delivery-manager-api/internal/api/http"
	"github.com/foliveiracamara/delivery-manager-api/internal/application/usecase"
	"github.com/foliveiracamara/delivery-manager-api/internal/domain"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/config"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/integration"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/persistence"
	"github.com/foliveiracamara/delivery-manager-api/internal/service"
	"go.uber.org/fx"
)

func providers() []any {
	globalDeps := []any{
		config.New,

		// Repository
		fx.Annotate(persistence.NewInMemoryPackageRepository, fx.As(new(domain.PackageRepository))),

		// Carrier Repository (MOCK)
		integration.NewCarrierRepository,

		// Services
		service.NewPackageService,

		// Use Cases
		ProvidePackageUseCase,

		// HTTP
		http.NewControllerManager,
	}

	globalDeps = append(globalDeps, http.ControllersList...)
	return globalDeps
}

func ProvidePackageUseCase() *usecase.PackageUseCase {
	repository := persistence.NewInMemoryPackageRepository()
	carrierRepo := integration.NewCarrierRepository()
	service := service.NewPackageService(carrierRepo)
	return usecase.NewPackage(repository, service)
}
