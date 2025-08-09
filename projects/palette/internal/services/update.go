package services

import "hamroh/palette/internal/models"

func (s *service) Update(palette *models.Palette) error {
	return s.repository.Update(palette)
}
