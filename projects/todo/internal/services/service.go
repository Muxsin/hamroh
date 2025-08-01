package services

import "hamroh-todo/internal/models"

type TodoRepository interface {
	Insert(todo *models.Todo) error
	All() ([]models.Todo, error)
}

type Service struct {
	repository TodoRepository
}

func New(repository TodoRepository) *Service {
	return &Service{
		repository: repository,
	}
}
