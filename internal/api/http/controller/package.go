package controller

import (
	"net/http"

	"github.com/foliveiracamara/delivery-manager-api/internal/api/http/dto"
	"github.com/foliveiracamara/delivery-manager-api/internal/application/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PackageController struct {
	us        *usecase.PackageUseCase
	validator *validator.Validate
}

func NewPackageController(usecase *usecase.PackageUseCase) *PackageController {
	return &PackageController{
		us:        usecase,
		validator: validator.New(),
	}
}

func (c *PackageController) Create(ctx echo.Context) error {
	req := &dto.PackageRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}
	if err := c.validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}

	id, err := c.us.Create(*req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "Package created successfully",
		"id":      id,
	})
}

func (c *PackageController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	pkg, err := c.us.Get(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	res := dto.PackageResponse{
		ID:            pkg.ID,
		Product:       pkg.Product,
		WeightKg:      pkg.WeightKg,
		EstadoDestino: pkg.DestinationState,
		RegiaoDestino: string(pkg.DestinationRegion),
		Status:        string(pkg.Status),
	}

	if pkg.Shipping != nil {
		res.Shipping = &dto.ShippingQuoteResponse{
			CarrierID:     pkg.Shipping.CarrierID,
			EstimatedDays: pkg.Shipping.EstimatedDays,
			Price:         pkg.Shipping.EstimatedPrice,
			CarrierName:   pkg.Shipping.CarrierName,
		}
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *PackageController) UpdateStatus(ctx echo.Context) error {
	req := &dto.UpdateStatusRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid request body"})
	}
	if err := c.validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}

	err := c.us.UpdateStatus(req.PackageID, req.Status)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Status updated successfully",
	})
}

// TODO: Delete later
func (c *PackageController) GetAll(ctx echo.Context) error {
	pkgs, err := c.us.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, pkgs)
}

func (c *PackageController) QuoteShippings(ctx echo.Context) error {
	req := &dto.ShippingsQuoteRequest{}
	req.PackageID = ctx.Param("id")

	shippings, err := c.us.QuoteShipping(req.PackageID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	res := dto.ShippingsQuoteResponse{
		Shippings: []dto.ShippingQuoteResponse{},
	}

	for _, shipping := range shippings {
		res.Shippings = append(res.Shippings, dto.ShippingQuoteResponse{
			CarrierID:     shipping.CarrierID,
			EstimatedDays: shipping.EstimatedDays,
			Price:         shipping.EstimatedPrice,
			CarrierName:   shipping.CarrierName,
		})
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *PackageController) HireCarrier(ctx echo.Context) error {
	req := &dto.HireCarrierRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid request body"})
	}
	if err := c.validator.Struct(req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}

	err := c.us.HireCarrier(req.PackageID, req.CarrierID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Carrier hired successfully",
	})
}
