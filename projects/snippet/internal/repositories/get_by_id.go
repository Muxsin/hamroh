package repositories

import "hamroh/snippet/internal/models"

func (r *repository) GetById(id string) (*models.Snippet, error) {
	var snippet models.Snippet
	if err := r.db.First(&snippet, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &snippet, nil
}
