package repositories

import "hamroh/flashcard/internal/models"

func (r *repository) GetAll() ([]*models.Flashcard, error) {
	var flashcards []*models.Flashcard

	if err := r.db.Find(&flashcards).Error; err != nil {
		return nil, err
	}

	return flashcards, nil
}
