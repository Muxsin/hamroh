package repositories

import "hamroh-todo/internal/models"

func (r *Repository) All() ([]models.Todo, error) {
	var todos []models.Todo

	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}
