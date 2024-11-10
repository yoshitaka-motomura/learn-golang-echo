package todos

import (
	"github.com/yoshitaka-motomura/learn-golang-echo/internal/todos/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *Repository) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}
func (r *Repository) CreateTodo(todo models.Todo) error {
	return r.db.Create(&todo).Error
}