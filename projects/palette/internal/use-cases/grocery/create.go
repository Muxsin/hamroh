package palette

import "hamroh/palette/internal/models"

func (uc *useCase) Create(palette *models.Palette) error {
	return uc.service.Create(palette)
}
