package repositories

import "hamroh/bookmark/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Bookmark{}, "id = ?", id).Error
}
