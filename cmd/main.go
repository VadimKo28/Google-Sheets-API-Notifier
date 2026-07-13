package main

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/lib/logger"
	"google_sheets_api/internal/server"
	"google_sheets_api/internal/service"
	"google_sheets_api/pkg/mail"
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
  
  client := mail.NewMailClient(
    cfg.GmailUser,   // username — для SMTP-аутентификации
    cfg.AppPassword,    // appPassword - "sayrymzexbixjvqa"
    cfg.GmailUser,   // from — что увидит получатель в поле "От кого"
	  log,
  )

  err = client.Send(
    cfg.GmailUser,
		"Данные из таблицы",
		"<h1>Отчёт</h1><p>Данные успешно обработаны.</p>",
	)

	if err != nil {
		log.Error("failed to send email: %v", "error", err)
	}

  server.Run()

}
