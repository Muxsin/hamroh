package repositories

import "hamroh/palette/internal/models"

func (r *repository) GetAll() ([]*models.Palette, error) {
	var palettes []*models.Palette

	if err := r.db.Find(&palettes).Error; err != nil {
		return nil, err
	}

	return palettes, nil
}
