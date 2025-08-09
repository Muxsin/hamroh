package repositories

import "hamroh/presentation/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Presentation{}, "id = ?", id).Error
}
