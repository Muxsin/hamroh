package repositories

import "hamroh/recipe/internal/models"

func (r *repository) Update(recipe *models.Recipe) error {
	return r.db.Updates(recipe).Error
}
