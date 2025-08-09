package services

import (
	"hamroh/presentation/internal/models"
)

func (s *service) Create(presentation *models.Presentation) error {
	return s.repository.Insert(presentation)
}
