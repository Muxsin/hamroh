package services

import (
	"kodnavis/people/internal/models"
)

func (s *service) Create(people *models.People) error {
	return s.repository.Insert(people)
}
