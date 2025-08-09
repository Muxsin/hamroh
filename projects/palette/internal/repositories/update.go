package repositories

import "hamroh/palette/internal/models"

func (r *repository) Update(palette *models.Palette) error {
	return r.db.Updates(palette).Error
}
