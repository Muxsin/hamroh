package repositories

import "hamroh/snippet/internal/models"

func (r *repository) Insert(snippet *models.Snippet) error {
	return r.db.Create(snippet).Error
}
