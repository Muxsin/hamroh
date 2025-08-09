package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type recipeResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	recipes, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []recipeResponse

	for _, recipe := range recipes {
		response = append(response, recipeResponse{
			Id:        recipe.ID,
			Text:      recipe.Text,
			CreatedAt: recipe.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
