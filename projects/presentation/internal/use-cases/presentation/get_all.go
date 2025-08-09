package presentation

import "hamroh/presentation/internal/models"

func (uc *useCase) GetAll() ([]*models.Presentation, error) {
	presentation, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return presentation, nil
}
