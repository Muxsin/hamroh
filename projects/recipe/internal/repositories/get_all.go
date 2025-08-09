package repositories

import "hamroh/recipe/internal/models"

func (r *repository) GetAll() ([]*models.Recipe, error) {
	var recipes []*models.Recipe

	if err := r.db.Find(&recipes).Error; err != nil {
		return nil, err
	}

	return recipes, nil
}
