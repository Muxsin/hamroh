package repositories

import "hamroh/recipe/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Recipe{}, "id = ?", id).Error
}
