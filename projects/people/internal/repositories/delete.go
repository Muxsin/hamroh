package repositories

import "kodnavis/people/internal/models"

func (r *repository) Delete(id string) error {
	return r.db.Delete(&models.People{}, "id = ?", id).Error
}
