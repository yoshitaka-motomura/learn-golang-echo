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
/**
 * NewTodoHandler creates a new TodoHandler
 * @param service *todos.Service
 * @return *TodoHandler
 * @description This function creates a new TodoHandler with the given service
 * @example
 * service := todos.NewService()
 * handler := NewTodoHandler(service)
 * @since 1.0.0
 * @version 1.0.0
 */
func NewTodoHandler(service *todos.Service) *TodoHandler {
	return &TodoHandler{service: service}
}

/**
 * GetTodos returns all todos
 * @param c echo.Context
 * @return error
 * @description This function returns all todos
 * @example
 * service := todos.NewService()
 * handler := NewTodoHandler(service)
 * e.GET("/todos", handler.GetTodos)
 * @since 1.0.0
 * @version 1.0.0
 */
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
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
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

func (h *TodoHandler) DeleteTodoByID(c echo.Context) error {
	id := c.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	err = h.service.DeleteTodoByID(uint(parsedID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}