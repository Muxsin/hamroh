package services

import "hamroh/tutorial/internal/models"

func (s *service) Update(tutorial *models.Tutorial) error {
	return s.repository.Update(tutorial)
}
