package palette

import "hamroh/palette/internal/models"

func (uc *useCase) GetOne(id string) (*models.Palette, error) {
	return uc.service.GetOne(id)
}
