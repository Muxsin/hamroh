package repositories

import "hamroh/grocery/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Grocery{}, "id = ?", id).Error
}
