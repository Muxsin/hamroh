package services

import (
	"kodnavis/password/internal/models"
)

func (s *service) Create(password *models.Password) error {
	return s.repository.Insert(password)
}
