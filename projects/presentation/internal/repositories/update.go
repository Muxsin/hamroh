package repositories

import "hamroh/presentation/internal/models"

func (r *repository) Update(presentation *models.Presentation) error {
	return r.db.Updates(presentation).Error
}
