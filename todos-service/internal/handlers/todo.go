package handlers

import (
	"github.com/gin-gonic/gin"
	"todos-service/internal/repositories"
)

type TodoHandlerInterface interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	GetById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TodoHandler struct {
	todo_repository repositories.TodoRepositoryInterface
}

func NewTodoHandler(todo_repository repositories.TodoRepositoryInterface) TodoHandlerInterface {
	return &TodoHandler{todo_repository: todo_repository}
}

func (h *TodoHandler) Create(ctx *gin.Context) {}

func (h *TodoHandler) List(ctx *gin.Context) {}

func (h *TodoHandler) GetById(ctx *gin.Context) {}

func (h *TodoHandler) Update(ctx *gin.Context) {}

func (h *TodoHandler) Delete(ctx *gin.Context) {}
