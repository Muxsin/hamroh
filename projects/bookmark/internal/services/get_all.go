package services

import "hamroh/bookmark/internal/models"

func (s *service) GetAll() ([]*models.Bookmark, error) {
	bookmark, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return bookmark, nil
}
