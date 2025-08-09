package services

import "hamroh/note/internal/models"

func (s *service) GetAll() ([]*models.Note, error) {
	note, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return note, nil
}
