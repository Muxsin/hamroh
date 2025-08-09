package http

import (
	"hamroh/presentation/internal/models"

	"github.com/go-playground/validator/v10"
)

type presentationUseCase interface {
	Create(presentation *models.Presentation) error
	GetAll() ([]*models.Presentation, error)
	GetOne(id string) (*models.Presentation, error)
	Delete(id string) error
	Update(presentation *models.Presentation) error
}

type handler struct {
	useCase presentationUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase presentationUseCase) *handler {
	return &handler{useCase: useCase}
}
