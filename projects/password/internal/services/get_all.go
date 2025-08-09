package services

import "kodnavis/password/internal/models"

func (s *service) GetAll() ([]*models.Password, error) {
	password, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return password, nil
}
