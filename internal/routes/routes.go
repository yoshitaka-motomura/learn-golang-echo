package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
)


func SetupRoutes(e *echo.Echo) {
	generalGroup := e.Group("")
	generalGroup.GET("/ping", func(c echo.Context) error {
		logging.Logger().Info("Ping endpoint called")
		return c.String(http.StatusOK, "pong")
	})

	generalGroup.GET("/hello", func(c echo.Context) error {
		logging.Logger().Info("Hello endpoint called")
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, world", "path": c.Path()})
	})
	TodosRoutes(e.Group("/todos"))
}