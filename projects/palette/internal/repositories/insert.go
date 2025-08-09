package repositories

import "hamroh/palette/internal/models"

func (r *repository) Insert(palette *models.Palette) error {
	return r.db.Create(palette).Error
}
