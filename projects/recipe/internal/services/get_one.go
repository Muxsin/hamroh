package services

import "hamroh/recipe/internal/models"

func (s *service) GetOne(id string) (*models.Recipe, error) {
	return s.repository.GetById(id)
}
