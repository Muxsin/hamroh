package services

import "hamroh/flashcard/internal/models"

func (s *service) GetAll() ([]*models.Flashcard, error) {
	flashcard, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return flashcard, nil
}
