package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/database"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/handlers"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/services/todos"
)

func TodosRoutes(group *echo.Group) {
	// GORMのDBインスタンスを取得
	db := database.DB

	// RepositoryとServiceを初期化してHandlerに渡す
	repo := todos.NewRepository(db)
	service := todos.NewService(repo)
	handler := handlers.NewTodoHandler(service)

	// ルーティングの設定
	group.GET("", handler.GetTodos)
	group.GET("/:id", handler.GetTodoById)
	group.POST("", handler.CreateTodo)
	// group.POST("", handler.CreateTodo)
	// group.PUT("/:id", handler.UpdateTodo)
	// group.DELETE("/:id", handler.DeleteTodo)
}
