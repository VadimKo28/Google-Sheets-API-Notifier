package server

import (
	"google_sheets_api/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct{
	handler *handler.Handler
	router *gin.Engine
}

func New(handler *handler.Handler, router *gin.Engine) *Server {
  return &Server{
	handler: handler,
	router: router,
  }
}

func(srv *Server) Register() {
  srv.router.GET("/sync_sheets", srv.handler.SyncSheets)
}

func (srv *Server) Run() {
	srv.router.Run(":8080")
}