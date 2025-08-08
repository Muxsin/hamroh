package services

import "kodnavis/people/internal/models"

func (s *service) Update(people *models.People) error {
	return s.repository.Update(people)
}
