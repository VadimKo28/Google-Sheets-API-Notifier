package handler

import (
	"context"
	"google_sheets_api/internal/domain"
	"github.com/gin-gonic/gin"
)

type Handler struct {
  servise GoogleSheets
}

type GoogleSheets interface {
  GetAndMappingSheets(c context.Context)([]domain.Event, error)
}

func New(service GoogleSheets) *Handler {
	return &Handler{	
		servise: service,
	}
}

func (h *Handler) GetSheets(c *gin.Context) {
  event, _ := h.servise.GetAndMappingSheets(c.Request.Context())

  c.JSON(200, event)
}