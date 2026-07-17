package sheets

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	GoogleSheetsService  Syncer
}

type Syncer interface {
	SyncSheets(c context.Context) error
}

func NewHandler(googleService Syncer) *Handler {
	return &Handler{
		GoogleSheetsService:  googleService,
	}
}

func (h *Handler) SyncSheets(c *gin.Context) {
	err := h.GoogleSheetsService.SyncSheets(c.Request.Context())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, nil)
}
