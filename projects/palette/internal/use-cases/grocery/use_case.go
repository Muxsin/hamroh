package palette

import "hamroh/palette/internal/models"

type paletteService interface {
	Create(palette *models.Palette) error
	GetAll() ([]*models.Palette, error)
	GetOne(id string) (*models.Palette, error)
	Delete(id string) error
	Update(palette *models.Palette) error
}

type useCase struct {
	service paletteService
}

func New(service paletteService) *useCase {
	return &useCase{service: service}
}
