package repositories

import "kodnavis/people/internal/models"

func (r *repository) GetAll() ([]*models.People, error) {
	var peoples []*models.People

	if err := r.db.Find(&peoples).Error; err != nil {
		return nil, err
	}

	return peoples, nil
}
