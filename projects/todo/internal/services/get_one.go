package services

import "hamroh-todo/internal/models"

func (s *service) GetOne(id string) (models.Todo, error) {
	todo, err := s.repository.GetById(id)
	if err != nil {
		return todo, err
	}

	return todo, nil
}
