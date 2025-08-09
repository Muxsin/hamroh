package presentation

import "hamroh/presentation/internal/models"

func (uc *useCase) Update(presentation *models.Presentation) error {
	return uc.service.Update(presentation)
}
