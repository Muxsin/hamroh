package services

import "hamroh/bookmark/internal/models"

func (s *service) GetOne(id string) (*models.Bookmark, error) {
	return s.repository.GetById(id)
}
