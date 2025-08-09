package repositories

import "hamroh/grocery/internal/models"

func (r *repository) Update(grocery *models.Grocery) error {
	return r.db.Updates(grocery).Error
}
