package services

import "hamroh/presentation/internal/models"

func (s *service) GetOne(id string) (*models.Presentation, error) {
	return s.repository.GetById(id)
}
