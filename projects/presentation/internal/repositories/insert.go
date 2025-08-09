package repositories

import "hamroh/presentation/internal/models"

func (r *repository) Insert(presentation *models.Presentation) error {
	return r.db.Create(presentation).Error
}
