package services

import "hamroh/flashcard/internal/models"

func (s *service) GetOne(id string) (*models.Flashcard, error) {
	return s.repository.GetById(id)
}
