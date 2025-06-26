package middlewares

import (
	"net/http"

	apperr "github.com/foliveiracamara/delivery-manager-api/internal/shared/apperror"
	"github.com/labstack/echo/v4"
)

// ErrorHandler é um middleware que trata AppErr e retorna os status codes HTTP apropriados
func ErrorHandler(err error, c echo.Context) {
	if appErr, ok := err.(*apperr.AppErr); ok {
		c.JSON(appErr.Code, appErr)
		return
	}

	c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal server error",
	})
}
