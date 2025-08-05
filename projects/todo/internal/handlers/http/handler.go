package http

import (
	"hamroh-todo/internal/models"
)

type useCase interface {
	Create(title string) (*models.Todo, error)
	List() ([]models.Todo, error)
	GetById(id string) (models.Todo, error)
	Update(todo *models.Todo) error
	Delete(id string) error
}

type handler struct {
	useCase useCase
}

func New(useCase useCase) *handler {
	return &handler{
		useCase: useCase,
	}
}
