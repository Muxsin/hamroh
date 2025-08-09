package services

import "hamroh/tutorial/internal/models"

func (s *service) GetOne(id string) (*models.Tutorial, error) {
	return s.repository.GetById(id)
}
