package http

import "github.com/labstack/echo/v4/middleware"

func (s *Server) setupMiddlewares() {

	// s.e.Pre(middleware.AddTrailingSlash())
	// s.e.Use(middleware.RequestID())
	// s.e.Use(middlewares.NewAuthMiddleware(authSvc))
	s.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// s.e.Use(middleware.Gzip())
	// s.e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Println("request", c.Response().Header().Get(echo.HeaderXRequestID), "path", c.Request().URL, "body", string(reqBody))
	// 	fmt.Println("response", c.Response().Header().Get(echo.HeaderXRequestID), "status", c.Response().Status, "body", string(resBody))
	// }))
	s.e.Use(middleware.Recover())
}
