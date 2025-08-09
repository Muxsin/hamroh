package services

import "hamroh/flashcard/internal/models"

func (s *service) Update(flashcard *models.Flashcard) error {
	return s.repository.Update(flashcard)
}
