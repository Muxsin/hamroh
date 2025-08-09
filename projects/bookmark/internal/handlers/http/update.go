package http

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateBookmarkRequest struct {
	Text string `json:"text" validate:"required"`
}

type updateBookmarkResponse struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

func (h *handler) Update(c *gin.Context) {
	id := c.Param("id")
	bookmark, err := h.useCase.GetOne(id)
	if err != nil {
		log.Println(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	var req updateBookmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookmark.Text = req.Text

	if err := h.useCase.Update(bookmark); err != nil {
		log.Println(err)

		c.Status(http.StatusInternalServerError)
		return
	}

	response := updateBookmarkResponse{
		Id:        id,
		Text:      bookmark.Text,
		CreatedAt: bookmark.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
