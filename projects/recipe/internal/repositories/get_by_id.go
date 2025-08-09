package repositories

import "hamroh/recipe/internal/models"

func (r *repository) GetById(id string) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := r.db.First(&recipe, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &recipe, nil
}
