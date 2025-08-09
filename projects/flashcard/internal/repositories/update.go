package repositories

import "hamroh/flashcard/internal/models"

func (r *repository) Update(flashcard *models.Flashcard) error {
	return r.db.Updates(flashcard).Error
}
