package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/handlers"
)

func TodosRoutes(group *echo.Group) {
	group.GET("", handlers.GetTodos)
	group.GET("/:id", handlers.GetTodoById)
}