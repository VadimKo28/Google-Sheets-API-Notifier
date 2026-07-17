package server

import (
	"google_sheets_api/internal/handler/events"
	"google_sheets_api/internal/handler/sheets"

	"github.com/gin-gonic/gin"
)

type Server struct {
	eventsHandler *events.Handler
	sheetsHandler *sheets.Handler
	router  *gin.Engine
}

func New(sheetsHandler *sheets.Handler,  eventsHandler *events.Handler, router *gin.Engine) *Server {
	return &Server{
		sheetsHandler: sheetsHandler,
		eventsHandler: eventsHandler,
		router:  router,
	}
}

func (srv *Server) Register() {
	srv.router.POST("/sync_sheets", srv.sheetsHandler.SyncSheets)
	srv.router.POST("/check_today_events", srv.eventsHandler.CheckEvents)
}

func (srv *Server) Run() {
	srv.router.Run(":8080")
}
