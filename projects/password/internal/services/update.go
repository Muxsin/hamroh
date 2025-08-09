package services

import "kodnavis/password/internal/models"

func (s *service) Update(password *models.Password) error {
	return s.repository.Update(password)
}
