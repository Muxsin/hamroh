package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bookmarkResponse struct {
	Id        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) List(c *gin.Context) {
	bookmarks, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	var response []bookmarkResponse

	for _, bookmark := range bookmarks {
		response = append(response, bookmarkResponse{
			Id:        bookmark.ID,
			Text:      bookmark.Text,
			CreatedAt: bookmark.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
