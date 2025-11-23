package entrypoint

import "github.com/labstack/echo/v4"

type Handler interface {
	RegisterRoutes(e *echo.Echo)
}

func RegisterRoute(e *echo.Echo, handler ...Handler) {
	for _, route := range handler {
		route.RegisterRoutes(e)
	}
}
