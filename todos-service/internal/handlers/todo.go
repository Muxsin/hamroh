package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"todos-service/internal/dto/requests"
	"todos-service/internal/dto/responses"
	"todos-service/internal/models"
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

func (h *TodoHandler) Create(ctx *gin.Context) {
	var request requests.CreateTodoRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := &models.Todo{
		Text:   request.Text,
		Done:   request.Done,
		UserID: request.UserID,
	}

	if err := h.todo_repository.Insert(todo); err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	response := responses.TodoResponse{
		Id:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		UserID:    todo.UserID,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) List(ctx *gin.Context) {
	todos, err := h.todo_repository.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todos": todos})
}

func (h *TodoHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := h.todo_repository.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	response := responses.TodoResponse{
		Id:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		UserID:    todo.UserID,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *TodoHandler) Update(ctx *gin.Context) {}

func (h *TodoHandler) Delete(ctx *gin.Context) {}
