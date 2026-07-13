package main

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/lib/logger"
	"google_sheets_api/internal/server"
	"google_sheets_api/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	log := logger.LoggerSetup()

    ctx := context.Background()
	spreadsheetId := cfg.SpreadsSheetID
	readRange := cfg.ReadRange

	service, err := service.NewGoogleSheetsService(ctx, spreadsheetId, readRange, log)

	if err != nil {
		log.Error("error create service", slog.Any("error", err))
		return
	}

	handler := handler.New(service)
	server := server.New(handler, gin.New())
	server.Register()
	server.Run()
}
