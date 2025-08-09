package services

import (
	"hamroh/snippet/internal/models"
)

func (s *service) Create(snippet *models.Snippet) error {
	return s.repository.Insert(snippet)
}
