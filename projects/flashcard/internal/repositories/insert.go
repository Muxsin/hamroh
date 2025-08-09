package repositories

import "hamroh/flashcard/internal/models"

func (r *repository) Insert(flashcard *models.Flashcard) error {
	return r.db.Create(flashcard).Error
}
