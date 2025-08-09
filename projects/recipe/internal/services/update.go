package services

import "hamroh/recipe/internal/models"

func (s *service) Update(recipe *models.Recipe) error {
	return s.repository.Update(recipe)
}
