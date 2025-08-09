package services

import (
	"hamroh/tutorial/internal/models"
)

func (s *service) Create(tutorial *models.Tutorial) error {
	return s.repository.Insert(tutorial)
}
