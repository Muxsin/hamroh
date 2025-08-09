package http

import (
	"hamroh/palette/internal/models"

	"github.com/go-playground/validator/v10"
)

type paletteUseCase interface {
	Create(palette *models.Palette) error
	GetAll() ([]*models.Palette, error)
	GetOne(id string) (*models.Palette, error)
	Delete(id string) error
	Update(palette *models.Palette) error
}

type handler struct {
	useCase paletteUseCase
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func New(useCase paletteUseCase) *handler {
	return &handler{useCase: useCase}
}
