package services

import "hamroh/snippet/internal/models"

func (s *service) GetOne(id string) (*models.Snippet, error) {
	return s.repository.GetById(id)
}
