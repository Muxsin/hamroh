package use_cases

import "hamroh-todo/internal/models"

type todoService interface {
	Create(todo *models.Todo) error
	List() ([]models.Todo, error)
	GetById(id string) (models.Todo, error)
}

type useCase struct {
	service todoService
}

func New(service todoService) *useCase {
	return &useCase{
		service: service,
	}
}
