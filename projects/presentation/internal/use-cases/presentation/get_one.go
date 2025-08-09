package presentation

import "hamroh/presentation/internal/models"

func (uc *useCase) GetOne(id string) (*models.Presentation, error) {
	return uc.service.GetOne(id)
}
