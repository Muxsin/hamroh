package services

import "hamroh/presentation/internal/models"

func (s *service) Update(presentation *models.Presentation) error {
	return s.repository.Update(presentation)
}
