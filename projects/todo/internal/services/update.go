package services

import "hamroh-todo/internal/models"

func (s *service) Update(todo *models.Todo) error {
	return s.repository.Update(todo)
}
