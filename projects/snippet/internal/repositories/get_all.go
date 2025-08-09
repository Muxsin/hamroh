package repositories

import "hamroh/snippet/internal/models"

func (r *repository) GetAll() ([]*models.Snippet, error) {
	var snippets []*models.Snippet

	if err := r.db.Find(&snippets).Error; err != nil {
		return nil, err
	}

	return snippets, nil
}
