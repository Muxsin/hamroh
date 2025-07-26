package responses

type TodoResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"Text"`
	Done      bool   `json:"body"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
