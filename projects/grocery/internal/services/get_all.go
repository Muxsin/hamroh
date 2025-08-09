package services

import "hamroh/grocery/internal/models"

func (s *service) GetAll() ([]*models.Grocery, error) {
	grocery, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return grocery, nil
}
