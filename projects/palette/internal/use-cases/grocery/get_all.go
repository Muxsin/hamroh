package palette

import "hamroh/palette/internal/models"

func (uc *useCase) GetAll() ([]*models.Palette, error) {
	palette, err := uc.service.GetAll()
	if err != nil {
		return nil, err
	}

	return palette, nil
}
