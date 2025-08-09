package repositories

import "hamroh/flashcard/internal/models"

func (r *repository) GetById(id string) (*models.Flashcard, error) {
	var flashcard models.Flashcard
	if err := r.db.First(&flashcard, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &flashcard, nil
}
