package services

import "kodnavis/people/internal/models"

func (s *service) GetOne(id string) (*models.People, error) {
	return s.repository.GetById(id)
}
