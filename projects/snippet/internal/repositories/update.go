package repositories

import "hamroh/snippet/internal/models"

func (r *repository) Update(snippet *models.Snippet) error {
	return r.db.Updates(snippet).Error
}
