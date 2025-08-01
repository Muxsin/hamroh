package repositories

import "hamroh-todo/internal/models"

func (r *repository) GetById(id string) (models.Todo, error) {
	var todo models.Todo

	result := r.db.Where("id = ?", id).First(&todo)
	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}
