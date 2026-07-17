package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	GoogleSheetsService  GoogleSheetsSyncer
	EvventNotifierSrvice EventNotifier
}

type EventNotifier interface {
	CheckEventsToday(c context.Context) error
}

type GoogleSheetsSyncer interface {
	SyncSheets(c context.Context) error
}

func New(googleService GoogleSheetsSyncer, eventService EventNotifier) *Handler {
	return &Handler{
		GoogleSheetsService:  googleService,
		EvventNotifierSrvice: eventService,
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

func (h *Handler) CheckEvents(c *gin.Context) {
	err := h.EvventNotifierSrvice.CheckEventsToday(c.Request.Context())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, nil)
}
