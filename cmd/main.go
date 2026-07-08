package main

import (
	"context"
	"fmt"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/handler"
	"google_sheets_api/internal/server"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
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

	handler := handler.New()
	server := server.New(handler, gin.Default())
	server.Register()
	server.Run()
	

	srv, err := sheets.NewService(ctx, option.WithAuthCredentialsFile(option.ServiceAccount, "my-sheets-integration-501715-8e3105270262.json"))
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			for _, cell := range row {
				fmt.Printf("%v\t", cell)
			}
			fmt.Println()
		}
	}
}