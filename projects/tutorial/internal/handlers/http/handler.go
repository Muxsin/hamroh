package http

import (
	"hamroh/tutorial/internal/models"

	"github.com/go-playground/validator/v10"
)

type tutorialUseCase interface {
	Create(tutorial *models.Tutorial) error
	GetAll() ([]*models.Tutorial, error)
	GetOne(id string) (*models.Tutorial, error)
	Delete(id string) error
	Update(tutorial *models.Tutorial) error
}

type handler struct {
	useCase tutorialUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase tutorialUseCase) *handler {
	return &handler{useCase: useCase}
}
