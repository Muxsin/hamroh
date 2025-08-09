package services

import "hamroh/palette/internal/models"

func (s *service) GetAll() ([]*models.Palette, error) {
	palette, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return palette, nil
}
