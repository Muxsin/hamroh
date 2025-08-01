package services

import "hamroh-todo/internal/models"

func (s *Service) List() ([]models.Todo, error) {
	todos, err := s.repository.All()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
