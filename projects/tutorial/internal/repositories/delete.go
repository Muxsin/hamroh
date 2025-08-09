package repositories

import "hamroh/tutorial/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Tutorial{}, "id = ?", id).Error
}
