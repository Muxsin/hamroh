package repositories

import "kodnavis/people/internal/models"

func (r *repository) Update(people *models.People) error {
	return r.db.Updates(people).Error
}
