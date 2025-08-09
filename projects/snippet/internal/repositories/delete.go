package repositories

import "hamroh/snippet/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Snippet{}, "id = ?", id).Error
}
