package repositories

import "kodnavis/people/internal/models"

func (r *repository) Insert(people *models.People) error {
	return r.db.Create(people).Error
}
