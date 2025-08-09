package repositories

import "hamroh/presentation/internal/models"

func (r *repository) GetById(id string) (*models.Presentation, error) {
	var presentation models.Presentation
	if err := r.db.First(&presentation, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &presentation, nil
}
