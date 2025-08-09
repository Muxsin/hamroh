package repositories

import "hamroh/grocery/internal/models"

func (r *repository) GetAll() ([]*models.Grocery, error) {
	var grocerys []*models.Grocery

	if err := r.db.Find(&grocerys).Error; err != nil {
		return nil, err
	}

	return grocerys, nil
}
