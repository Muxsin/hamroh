package repositories

import "hamroh/palette/internal/models"

func (r *repository) GetById(id string) (*models.Palette, error) {
	var palette models.Palette
	if err := r.db.First(&palette, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &palette, nil
}
