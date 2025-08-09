package services

import (
	"hamroh/recipe/internal/models"
)

func (s *service) Create(recipe *models.Recipe) error {
	return s.repository.Insert(recipe)
}
