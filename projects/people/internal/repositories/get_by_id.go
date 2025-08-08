package repositories

import "kodnavis/people/internal/models"

func (r *repository) GetById(id string) (*models.People, error) {
	var people models.People
	if err := r.db.First(&people, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &people, nil
}
