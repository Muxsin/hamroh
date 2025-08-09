package services

import (
	"hamroh/flashcard/internal/models"
)

func (s *service) Create(flashcard *models.Flashcard) error {
	return s.repository.Insert(flashcard)
}
