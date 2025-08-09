package repositories

import "hamroh/bookmark/internal/models"

func (r *repository) Update(bookmark *models.Bookmark) error {
	return r.db.Updates(bookmark).Error
}
