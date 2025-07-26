package repositories

import (
	"gorm.io/gorm"
	"todos-service/internal/models"
)

type TodoRepositoryInterface interface {
	Insert(todo models.Todo) error
	List() ([]models.Todo, error)
	GetById(id string) (models.Todo, error)
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

func (r *TodoRepository) Insert(todo models.Todo) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(todo).Error; err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func (r *TodoRepository) List() ([]models.Todo, error) {
	var todos []models.Todo

	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func (r *TodoRepository) GetById(id string) (models.Todo, error) {
	var todo models.Todo

	result := r.db.First(&todo, id)
	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

func (r *TodoRepository) Update(id string) error {
	return nil
}

func (r *TodoRepository) Delete(id string) error {
	return nil
}
