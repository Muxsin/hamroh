package repositories

import "hamroh/tutorial/internal/models"

func (r *repository) GetById(id string) (*models.Tutorial, error) {
	var tutorial models.Tutorial
	if err := r.db.First(&tutorial, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &tutorial, nil
}
