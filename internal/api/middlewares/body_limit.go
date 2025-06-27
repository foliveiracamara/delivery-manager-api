package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BodyLimitMiddleware limita o tamanho do body das requests
func BodyLimitMiddleware() echo.MiddlewareFunc {
	return middleware.BodyLimit("1MB")
}