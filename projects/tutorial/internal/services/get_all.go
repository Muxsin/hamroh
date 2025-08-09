package services

import "hamroh/tutorial/internal/models"

func (s *service) GetAll() ([]*models.Tutorial, error) {
	tutorial, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return tutorial, nil
}
