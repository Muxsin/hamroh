package repositories

import "kodnavis/password/internal/models"

func (r *repository) GetById(id string) (*models.Password, error) {
	var password models.Password
	if err := r.db.First(&password, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &password, nil
}
