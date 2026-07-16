package main

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/lib/logger"
	"google_sheets_api/internal/repository/event"
	"google_sheets_api/internal/server"
	"google_sheets_api/internal/service/event_notifier"
	"google_sheets_api/internal/service/google_sheets"
	"google_sheets_api/pkg/clients/gmail"
	sheets_client "google_sheets_api/pkg/clients/google_sheets"
	"google_sheets_api/pkg/clients/postgres"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

  spreadsheetId := cfg.SpreadsSheetID
	readRange := cfg.ReadRange

	log := logger.LoggerSetup()

  sheetsClient, err := sheets_client.NewClient(context.Background(), spreadsheetId, readRange, log)

  if err != nil {
    log.Error("error create service", slog.Any("error", err))
    return
  }

	dbClient, err := postgres.NewClient(context.Background(), cfg.PostgresConnStr)
  mailClient := gmail.NewClient(
    cfg.GmailUser,  
    cfg.AppPassword, 
    cfg.GmailUser,
	  log,
  )

  repository := event.NewEventRepository(dbClient, log)

	if err != nil {
	  log.Error("error connect to postgres", slog.Any("error", err))
	  return
	} else {
	  log.Info("connect to postgres")
	}

	googleSheetsService, err := google_sheets.NewGoogleSheetsService(sheetsClient, log, repository)
  eventNotifierService := event_notifier.NewEventNotifierService(log, repository, mailClient)

	if err != nil {
		log.Error("error create service", slog.Any("error", err))
		return
	}

	handler := handler.New(googleSheetsService, eventNotifierService)
	server := server.New(handler, gin.New())
	server.Register()
  
  server.Run()
}
