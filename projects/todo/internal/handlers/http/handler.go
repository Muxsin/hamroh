package http

import (
	"hamroh-todo/internal/models"
)

type UseCase interface {
	Create(title string) (*models.Todo, error)
	List() ([]models.Todo, error)
}

type Handler struct {
	UseCase UseCase
}

func New(useCase UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}
