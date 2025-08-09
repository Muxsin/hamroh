package services

import "hamroh/presentation/internal/models"

func (s *service) GetAll() ([]*models.Presentation, error) {
	presentation, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return presentation, nil
}
