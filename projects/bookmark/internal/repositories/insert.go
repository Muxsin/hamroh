package repositories

import "hamroh/bookmark/internal/models"

func (r *repository) Insert(bookmark *models.Bookmark) error {
	return r.db.Create(bookmark).Error
}
