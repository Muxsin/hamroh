package services

import "kodnavis/password/internal/models"

func (s *service) GetOne(id string) (*models.Password, error) {
	return s.repository.GetById(id)
}
