package http

import (
	"hamroh-todo/internal/models"
)

type useCase interface {
	Create(title string) (*models.Todo, error)
	List() ([]models.Todo, error)
}

type handler struct {
	useCase useCase
}

func New(useCase useCase) *handler {
	return &handler{
		useCase: useCase,
	}
}
