package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
  servise GoogleSheets
}

type GoogleSheets interface {
  SyncSheets(c context.Context)( error)
}

func New(service GoogleSheets) *Handler {
	return &Handler{	
		servise: service,
	}
}

func (h *Handler) SyncSheets(c *gin.Context) {
  err := h.servise.SyncSheets(c.Request.Context())

  if err != nil {
	c.JSON(500, err)
	return
  }

  c.JSON(200, nil)
}