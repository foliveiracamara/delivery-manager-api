package http

import (
	"context"
	"fmt"
	"time"

	"github.com/foliveiracamara/delivery-manager-api/internal/api"
	"github.com/foliveiracamara/delivery-manager-api/internal/api/middlewares"
	"github.com/foliveiracamara/delivery-manager-api/internal/infrastructure/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e *echo.Echo
}

var _ api.Api = (*Server)(nil)

func New(cfg *config.Config, controllerManager *ControllerManager) *Server {
	server := &Server{
		e: echo.New(),
	}

	server.setupMiddlewares()
	server.setupRoutes(controllerManager)

	server.e.Server.Addr = fmt.Sprintf(":%d", cfg.Server.Port)

	// server.e.HidePort = true
	server.e.HideBanner = true
	return server
}

func (s *Server) Run(ctx context.Context) error {
	return s.e.Start(s.e.Server.Addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) setupMiddlewares() {
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	s.e.Use(middleware.Recover())

	// Middlewares de segurança
	s.e.Use(middlewares.SecurityHeaders())
	s.e.Use(middlewares.CORSMiddleware())
	s.e.Use(middlewares.BodyLimitMiddleware())
	s.e.Use(middlewares.TimeoutMiddleware())
	s.e.Use(middlewares.RateLimitMiddleware(50, time.Minute))

	s.e.HTTPErrorHandler = middlewares.ErrorHandler
}
