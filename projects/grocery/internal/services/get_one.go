package services

import "hamroh/grocery/internal/models"

func (s *service) GetOne(id string) (*models.Grocery, error) {
	return s.repository.GetById(id)
}
