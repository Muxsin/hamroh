package http

import (
	"kodnavis/password/internal/models"

	"github.com/go-playground/validator/v10"
)

type passwordUseCase interface {
	Create(password *models.Password) error
	GetAll() ([]*models.Password, error)
	GetOne(id string) (*models.Password, error)
	Delete(id string) error
	Update(password *models.Password) error
}

type handler struct {
	useCase passwordUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase passwordUseCase) *handler {
	return &handler{useCase: useCase}
}
