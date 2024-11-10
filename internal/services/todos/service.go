package todos

import (
	"errors"

	"github.com/yoshitaka-motomura/learn-golang-echo/utils"

	"github.com/yoshitaka-motomura/learn-golang-echo/internal/models"
	"gorm.io/gorm"
)

// Serviceはtodosのビジネスロジックを管理する
type Service struct {
	repo *Repository
}

// NewServiceは新しいServiceインスタンスを返します
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetTodosは全てのTodoリストを取得します
func (s *Service) GetTodos() ([]models.Todo, error) {
	return s.repo.GetTodos()
}


func (s *Service) GetTodoByID(id uint) (*models.Todo, error) {
	todo, err := s.repo.GetTodoByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return todo, nil
}

func (s *Service) CreateTodo(todo models.Todo) (*models.Todo, error) {
	// バリデーションを実行
	if err := utils.ValidateStruct(todo); err != nil {
		return nil, err
	}

	// リポジトリでTodoを作成
	if err := s.repo.CreateTodo(todo); err != nil {
		return nil, err
	}
	return &todo, nil
}