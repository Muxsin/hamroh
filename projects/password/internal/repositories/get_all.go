package repositories

import "kodnavis/password/internal/models"

func (r *repository) GetAll() ([]*models.Password, error) {
	var passwords []*models.Password

	if err := r.db.Find(&passwords).Error; err != nil {
		return nil, err
	}

	return passwords, nil
}
