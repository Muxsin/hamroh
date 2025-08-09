package services

import "hamroh/palette/internal/models"

type paletteRepository interface {
	Insert(palette *models.Palette) error
	GetAll() ([]*models.Palette, error)
	GetById(id string) (*models.Palette, error)
	Delete(id string) error
	Update(palette *models.Palette) error
}

type service struct {
	repository paletteRepository
}

func New(repository paletteRepository) *service {
	return &service{
		repository: repository,
	}
}
