package http

import "github.com/foliveiracamara/delivery-manager-api/internal/api/http/controller"

type ControllerManager struct {
	PackageController *controller.PackageController
}

var ControllersList = []any{
	controller.NewPackageController,
}

func NewControllerManager(
	packageController *controller.PackageController,
) *ControllerManager {
	return &ControllerManager{
		PackageController: packageController,
	}
}
