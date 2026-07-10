package main

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/server"
	"google_sheets_api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
)

func main() {
	if err := godotenv.Load(); err != nil {
      log.Print("No .env file found")
    }

	cfg := config.NewConfig()
	log.Print("Config Load")

	spreadsheetId := cfg.SpreadsSheetID
	readRange := cfg.ReadRange

	ctx := context.Background()

	service := service.NewGoogleSheetsService(ctx, spreadsheetId, readRange)

	handler := handler.New(service)
	server := server.New(handler, gin.New())
	server.Register()
	server.Run()
}
