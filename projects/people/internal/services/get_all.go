package services

import "kodnavis/people/internal/models"

func (s *service) GetAll() ([]*models.People, error) {
	people, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return people, nil
}
