package main

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/lib/logger"
	"google_sheets_api/internal/server"
	"google_sheets_api/internal/service"
	"google_sheets_api/pkg/clients/gmail"
	"google_sheets_api/pkg/clients/postgres"
  "google_sheets_api/internal/repository/event"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	log := logger.LoggerSetup()

	db, err := postgres.NewClient(context.Background(), cfg.PostgresConnStr)

  repository := event.NewEventRepository(db)

	if err != nil {
	  log.Error("error connect to postgres", slog.Any("error", err))
	  return
	} else {
	  log.Info("connect to postgres")
	}

    ctx := context.Background()
	spreadsheetId := cfg.SpreadsSheetID
	readRange := cfg.ReadRange

	service, err := service.NewGoogleSheetsService(ctx, spreadsheetId, readRange, log, repository)

	if err != nil {
		log.Error("error create service", slog.Any("error", err))
		return
	}

	handler := handler.New(service)
	server := server.New(handler, gin.New())
	server.Register()
  
  client := gmail.NewClient(
    cfg.GmailUser,  
    cfg.AppPassword, 
    cfg.GmailUser,   // from — что увидит получатель в поле "От кого"
	  log,
  )

  err = client.Send(
    cfg.GmailUser,
		"Данные из таблицы",
		"<h1>Отчёт</h1><p>Данные успешно обработаны !!!!!.</p>",
	)

	if err != nil {
		log.Error("failed to send email: %v", "error", err)
	}

  server.Run()

}
