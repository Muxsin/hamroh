package services

import (
	"hamroh/grocery/internal/models"
)

func (s *service) Create(grocery *models.Grocery) error {
	return s.repository.Insert(grocery)
}
