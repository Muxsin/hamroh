package repositories

import (
	"gorm.io/gorm"
	"todos-service/internal/models"
)

type TodoRepositoryInterface interface {
	Insert(todo *models.Todo) error
	List() ([]*models.Todo, error)
	GetById(id string) (*models.Todo, error)
	Update(id string) error
	Delete(id string) error
}

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepositoryInterface {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Insert(todo *models.Todo) error {
	return nil
}

func (r *TodoRepository) List() ([]*models.Todo, error) {
	return nil, nil
}

func (r *TodoRepository) GetById(id string) (*models.Todo, error) {
	return nil, nil
}

func (r *TodoRepository) Update(id string) error {
	return nil
}

func (r *TodoRepository) Delete(id string) error {
	return nil
}
