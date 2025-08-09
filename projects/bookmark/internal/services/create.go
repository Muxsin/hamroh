package services

import (
	"hamroh/bookmark/internal/models"
)

func (s *service) Create(bookmark *models.Bookmark) error {
	return s.repository.Insert(bookmark)
}
