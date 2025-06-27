package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// HealthCheck godoc
// @Summary Health check da API
// @Description Verifica se a API está funcionando corretamente
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthCheckResponse
// @Router /health [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "OK!",
	})
}

func (s *Server) setupRoutes(cm *ControllerManager) {
	mainRouter := s.e.Group("")

	packageRouter := mainRouter.Group("/package")
	packageRouter.GET("/:id", cm.PackageController.Get)
	packageRouter.POST("/", cm.PackageController.Create)
	packageRouter.POST("/:id/quote", cm.PackageController.QuoteShippings)
	packageRouter.POST("/hire-carrier", cm.PackageController.HireCarrier)
	packageRouter.PUT("/status", cm.PackageController.UpdateStatus)

	mainRouter.GET("/health", healthCheck)

	// Swagger documentation
	mainRouter.GET("/swagger/*", echoSwagger.WrapHandler)
}
