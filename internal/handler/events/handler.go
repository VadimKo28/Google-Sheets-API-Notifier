package events

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	EvventNotifierSrvice Notifier
}

type Notifier interface {
	CheckEventsToday(c context.Context) error
}


func NewHandler(eventService Notifier) *Handler {
	return &Handler{
		EvventNotifierSrvice: eventService,
	}
}

func (h *Handler) CheckEvents(c *gin.Context) {
	err := h.EvventNotifierSrvice.CheckEventsToday(c.Request.Context())

	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, nil)
}
