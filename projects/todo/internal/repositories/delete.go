package repositories

import "hamroh-todo/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
