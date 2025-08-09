package repositories

import "hamroh/tutorial/internal/models"

func (r *repository) Update(tutorial *models.Tutorial) error {
	return r.db.Updates(tutorial).Error
}
