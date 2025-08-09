package repositories

import "kodnavis/password/internal/models"

func (r *repository) Insert(password *models.Password) error {
	return r.db.Create(password).Error
}
