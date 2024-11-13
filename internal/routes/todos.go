package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/database"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos/handlers"
)

func TodosRoutes(group *echo.Group) {
	db := database.DB

	repo := todos.NewRepository(db)
	service := todos.NewService(repo)
	handler := handlers.NewTodoHandler(service)

	// ルーティングの設定
	group.GET("", handler.GetTodos)
	group.GET("/:id", handler.GetTodoById)
	group.POST("", handler.CreateTodo)
	// group.PUT("/:id", handler.UpdateTodo)
	group.DELETE("/:id", handler.DeleteTodoByID)
}
