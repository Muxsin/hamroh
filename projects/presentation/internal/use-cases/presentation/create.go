package presentation

import "hamroh/presentation/internal/models"

func (uc *useCase) Create(presentation *models.Presentation) error {
	return uc.service.Create(presentation)
}
