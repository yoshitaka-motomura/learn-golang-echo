package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos"
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos/models"
	"github.com/yoshitaka-motomura/learn-golang-echo/utils"
)

type TodoHandler struct {
	service *todos.Service
}

func NewTodoHandler(service *todos.Service) *TodoHandler {
	return &TodoHandler{service: service}
}
func (h *TodoHandler) GetTodos(c echo.Context) error {
	todos, err := h.service.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodoById(c echo.Context) error {
	id := c.Param("id") // URLパラメータからIDを取得
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	todo, err := h.service.GetTodoByID(uint(parsedID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	_, err := h.service.CreateTodo(todo)
	if validationErrors, ok := err.(utils.ValidationErrors); ok {
		return c.JSON(http.StatusUnprocessableEntity, validationErrors)
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	

	return c.NoContent(http.StatusCreated)
}