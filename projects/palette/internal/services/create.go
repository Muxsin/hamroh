package services

import (
	"hamroh/palette/internal/models"
)

func (s *service) Create(palette *models.Palette) error {
	return s.repository.Insert(palette)
}
