package services

import "hamroh/palette/internal/models"

func (s *service) GetOne(id string) (*models.Palette, error) {
	return s.repository.GetById(id)
}
