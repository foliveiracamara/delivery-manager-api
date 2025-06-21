package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) setupRoutes(cm *ControllerManager) {
	mainRouter := s.e.Group("")

	packageRouter := mainRouter.Group("/package")
	packageRouter.GET("/:id", cm.PackageController.Get)
	packageRouter.GET("/", cm.PackageController.GetAll) //TODO: Delete later
	packageRouter.POST("/", cm.PackageController.Create)
	packageRouter.POST("/:id/quote", cm.PackageController.QuoteShippings)
	packageRouter.POST("/hire-carrier", cm.PackageController.HireCarrier)

	mainRouter.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "OK!",
		})
	})
}
