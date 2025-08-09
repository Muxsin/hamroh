package repositories

import "hamroh/note/internal/models"

func (r *repository) Update(note *models.Note) error {
	return r.db.Updates(note).Error
}
