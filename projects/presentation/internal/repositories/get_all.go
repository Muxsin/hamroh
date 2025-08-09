package repositories

import "hamroh/presentation/internal/models"

func (r *repository) GetAll() ([]*models.Presentation, error) {
	var presentations []*models.Presentation

	if err := r.db.Find(&presentations).Error; err != nil {
		return nil, err
	}

	return presentations, nil
}
