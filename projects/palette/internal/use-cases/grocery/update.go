package palette

import "hamroh/palette/internal/models"

func (uc *useCase) Update(palette *models.Palette) error {
	return uc.service.Update(palette)
}
