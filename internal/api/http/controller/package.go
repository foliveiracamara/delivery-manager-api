package controller

import (
	"net/http"

	"github.com/foliveiracamara/delivery-manager-api/internal/api/http/dto"
	"github.com/foliveiracamara/delivery-manager-api/internal/application/usecase"
	"github.com/labstack/echo/v4"
	// "github.com/go-playground/validator/v10"
)

type PackageController struct {
	us *usecase.PackageUseCase
	// validator *validator.Validate
}

func NewPackageController(usecase *usecase.PackageUseCase) *PackageController {
	return &PackageController{
		us: usecase,
	}
}

func (c *PackageController) Create(ctx echo.Context) error {
	dto := &dto.PackageRequest{}
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			map[string]string{"error": err.Error()})
	}

	err := c.us.Create(*dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated,
		map[string]string{"message": "Package created successfully"},
	)
}

func (c *PackageController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	pkg, err := c.us.Get(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]any{
		"id":                 pkg.ID,
		"product":            pkg.Product,
		"weight_kg":          pkg.WeightKg,
		"destination_region": pkg.DestinationRegion,
		"status":             pkg.Status,
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
	//Pegar id do pacote
	//Chamar usecase de cotação de frete
	//Retornar as cotações de frete, ordenadas por entrega mais rápida.
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

// func (c *PackageController) UpdateStatus(ctx echo.Context) error {}
