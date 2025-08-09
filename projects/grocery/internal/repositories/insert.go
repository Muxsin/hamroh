package repositories

import "hamroh/grocery/internal/models"

func (r *repository) Insert(grocery *models.Grocery) error {
	return r.db.Create(grocery).Error
}
