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

// Create godoc
// @Summary Criar um novo pacote
// @Description Cria um novo pacote com produto, peso e estado de destino. O sistema automaticamente mapeia o estado para a região correspondente e calcula as transportadoras disponíveis.
// @Tags packages
// @Accept json
// @Produce json
// @Param package body dto.PackageRequest true "Dados do pacote"
// @Success 201 {object} dto.CreatePackageResponse "Pacote criado com sucesso"
// @Router /package/ [post]
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
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "Package created successfully",
		"id":      id,
	})
}

// Get godoc
// @Summary Consultar um pacote específico
// @Description Retorna os dados completos de um pacote pelo ID, incluindo informações de entrega se uma transportadora foi contratada.
// @Tags packages
// @Accept json
// @Produce json
// @Param id path string true "ID único do pacote"
// @Success 200 {object} dto.PackageResponse "Dados do pacote"
// @Router /package/{id} [get]
func (c *PackageController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	pkg, err := c.us.Get(id)
	if err != nil {
		return err
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
			Transportadora:    pkg.Shipping.CarrierName,
			PrecoEstimado:     pkg.Shipping.EstimatedPrice,
			PrazoEstimadoDias: pkg.Shipping.EstimatedDays,
			CarrierID:         pkg.Shipping.CarrierID,
		}
	}

	return ctx.JSON(http.StatusOK, res)
}

// UpdateStatus godoc
// @Summary Atualizar status de um pacote
// @Description Atualiza o status de um pacote específico. Status válidos: criado, esperando_coleta, coletado, enviado, entregue, extraviado.
// @Tags packages
// @Accept json
// @Produce json
// @Param request body dto.UpdateStatusRequest true "Dados para atualização de status"
// @Success 200 {object} dto.SuccessResponse "Status atualizado com sucesso"
// @Router /package/status [put]
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
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Status updated successfully",
	})
}

// GetAll godoc
// @Summary Listar todos os pacotes
// @Description Retorna todos os pacotes cadastrados no sistema. Endpoint temporário para desenvolvimento.
// @Tags packages
// @Accept json
// @Produce json
// @Success 200 {array} dto.PackageResponse "Lista de todos os pacotes"
// @Router /package [get]
// TODO: Delete later
func (c *PackageController) GetAll(ctx echo.Context) error {
	pkgs, err := c.us.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, pkgs)
}

// QuoteShippings godoc
// @Summary Cotação de fretes
// @Description Retorna cotações de frete disponíveis para um pacote, ordenadas por prazo de entrega. Inclui preços e prazos estimados de todas as transportadoras que atendem a região do pacote.
// @Tags packages
// @Accept json
// @Produce json
// @Param id path string true "ID do pacote"
// @Success 200 {array} dto.ShippingQuoteResponse "Cotações de frete disponíveis"
// @Router /package/{id}/quote [post]
func (c *PackageController) QuoteShippings(ctx echo.Context) error {
	req := &dto.ShippingsQuoteRequest{}
	req.PackageID = ctx.Param("id")

	shippings, err := c.us.QuoteShipping(req.PackageID)
	if err != nil {
		return err
	}

	response := make([]dto.ShippingQuoteResponse, len(shippings))
	for i, shipping := range shippings {
		response[i] = dto.ShippingQuoteResponse{
			Transportadora:    shipping.CarrierName,
			PrecoEstimado:     shipping.EstimatedPrice,
			PrazoEstimadoDias: shipping.EstimatedDays,
			CarrierID:         shipping.CarrierID,
		}
	}

	return ctx.JSON(http.StatusOK, response)
}

// HireCarrier godoc
// @Summary Contratar transportadora
// @Description Contrata uma transportadora para realizar a entrega do pacote. O status do pacote será automaticamente alterado para 'esperando_coleta'. Transportadoras disponíveis: nebulix, rotafacil, moventra.
// @Tags packages
// @Accept json
// @Produce json
// @Param request body dto.HireCarrierRequest true "Dados para contratação"
// @Success 200 {object} dto.SuccessResponse "Transportadora contratada com sucesso"
// @Router /package/hire-carrier [post]
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
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Carrier hired successfully",
	})
}
