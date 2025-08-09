package repositories

import "hamroh/palette/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Palette{}, "id = ?", id).Error
}
