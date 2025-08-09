package repositories

import "hamroh/tutorial/internal/models"

func (r *repository) GetAll() ([]*models.Tutorial, error) {
	var tutorials []*models.Tutorial

	if err := r.db.Find(&tutorials).Error; err != nil {
		return nil, err
	}

	return tutorials, nil
}
