package repositories

import "hamroh/grocery/internal/models"

func (r *repository) GetById(id string) (*models.Grocery, error) {
	var grocery models.Grocery
	if err := r.db.First(&grocery, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &grocery, nil
}
