package services

import "hamroh/recipe/internal/models"

func (s *service) GetAll() ([]*models.Recipe, error) {
	recipe, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
