package services

import "hamroh/grocery/internal/models"

func (s *service) Update(grocery *models.Grocery) error {
	return s.repository.Update(grocery)
}
