package todo

import "hamroh-todo/internal/models"

type todoService interface {
	Create(todo *models.Todo) error
	GetAll() ([]models.Todo, error)
	GetOne(id string) (models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id string) error
}

type useCase struct {
	service todoService
}

func New(service todoService) *useCase {
	return &useCase{
		service: service,
	}
}
