package http

import "kodnavis/people/internal/models"

type peopleUseCase interface {
	Create(people *models.People) error
	GetAll() ([]*models.People, error)
	GetOne(id string) (*models.People, error)
}

type handler struct {
	useCase peopleUseCase
}

func New(useCase peopleUseCase) *handler {
	return &handler{useCase: useCase}
}
