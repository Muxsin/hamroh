package repositories

import "kodnavis/password/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.Password{}, "id = ?", id).Error
}
