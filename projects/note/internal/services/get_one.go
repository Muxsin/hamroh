package services

import "hamroh/note/internal/models"

func (s *service) GetOne(id string) (*models.Note, error) {
	return s.repository.GetById(id)
}
