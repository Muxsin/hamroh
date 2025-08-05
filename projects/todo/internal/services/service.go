package services

import "hamroh-todo/internal/models"

type todoRepository interface {
	Insert(todo *models.Todo) error
	All() ([]models.Todo, error)
	GetById(id string) (models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id string) error
}

type service struct {
	repository todoRepository
}

func New(repository todoRepository) *service {
	return &service{
		repository: repository,
	}
}
