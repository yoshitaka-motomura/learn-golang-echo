package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/logging"
)

func GetTodos(c echo.Context) error {
	logging.Logger().Info("GetTodos endpoint called")
	return c.JSON(http.StatusOK, []string{"todo1", "todo2", "todo3"})
}

func GetTodoById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}