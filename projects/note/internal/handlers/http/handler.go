package http

import (
	"hamroh/note/internal/models"

	"github.com/go-playground/validator/v10"
)

type noteUseCase interface {
	Create(note *models.Note) error
	GetAll() ([]*models.Note, error)
	GetOne(id string) (*models.Note, error)
	Delete(id string) error
	Update(note *models.Note) error
}

type handler struct {
	useCase noteUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase noteUseCase) *handler {
	return &handler{useCase: useCase}
}
