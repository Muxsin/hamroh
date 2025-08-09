package repositories

import "hamroh/tutorial/internal/models"

func (r *repository) Insert(tutorial *models.Tutorial) error {
	return r.db.Create(tutorial).Error
}
