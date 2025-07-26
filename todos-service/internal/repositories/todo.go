package repositories

import (
	"fmt"
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

var ErrFailedToCreate error

func (r *TodoRepository) Insert(todo *models.Todo) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(todo).Error; err != nil {
		tx.Rollback()
		ErrFailedToCreate = fmt.Errorf("failed to create user: %w", err)

		return ErrFailedToCreate
	}

	return tx.Commit().Error
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
