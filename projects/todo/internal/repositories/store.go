package repositories

import "hamroh-todo/internal/models"

func (r *repository) Insert(todo *models.Todo) error {
	return r.db.Create(todo).Error
}
