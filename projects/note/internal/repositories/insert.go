package repositories

import "hamroh/note/internal/models"

func (r *repository) Insert(note *models.Note) error {
	return r.db.Create(note).Error
}
