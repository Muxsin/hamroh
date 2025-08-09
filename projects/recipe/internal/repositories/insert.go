package repositories

import "hamroh/recipe/internal/models"

func (r *repository) Insert(recipe *models.Recipe) error {
	return r.db.Create(recipe).Error
}
