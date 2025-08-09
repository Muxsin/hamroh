package services

import "hamroh/snippet/internal/models"

func (s *service) GetAll() ([]*models.Snippet, error) {
	snippet, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return snippet, nil
}
