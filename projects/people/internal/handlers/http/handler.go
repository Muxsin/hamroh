package http

import (
	"github.com/go-playground/validator/v10"
	"kodnavis/people/internal/models"
)

type peopleUseCase interface {
	Create(people *models.People) error
	GetAll() ([]*models.People, error)
	GetOne(id string) (*models.People, error)
	Delete(id string) error
	Update(people *models.People) error
}

type handler struct {
	useCase peopleUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase peopleUseCase) *handler {
	return &handler{useCase: useCase}
}
