package repositories

import "kodnavis/password/internal/models"

func (r *repository) Update(password *models.Password) error {
	return r.db.Updates(password).Error
}
