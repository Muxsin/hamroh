package services

import "hamroh/snippet/internal/models"

func (s *service) Update(snippet *models.Snippet) error {
	return s.repository.Update(snippet)
}
