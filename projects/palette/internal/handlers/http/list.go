package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type paletteResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	palettes, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []paletteResponse

	for _, palette := range palettes {
		response = append(response, paletteResponse{
			Id:        palette.ID,
			Text:      palette.Text,
			CreatedAt: palette.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
