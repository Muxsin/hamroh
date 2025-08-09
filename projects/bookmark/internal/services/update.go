package services

import "hamroh/bookmark/internal/models"

func (s *service) Update(bookmark *models.Bookmark) error {
	return s.repository.Update(bookmark)
}
