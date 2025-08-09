package services

import "hamroh/note/internal/models"

func (s *service) Update(note *models.Note) error {
	return s.repository.Update(note)
}
