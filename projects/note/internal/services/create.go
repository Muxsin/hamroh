package services

import (
	"hamroh/note/internal/models"
)

func (s *service) Create(note *models.Note) error {
	return s.repository.Insert(note)
}
