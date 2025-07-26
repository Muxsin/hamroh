package requests

type UpdateTodoRequest struct {
	Text   string `json:"Text"`
	Done   bool   `json:"body"`
	UserID string `json:"user_id"`
}
