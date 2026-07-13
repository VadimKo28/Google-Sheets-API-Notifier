package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SpreadsSheetID string
  ReadRange string
}

func NewConfig() *Config {
  if err := godotenv.Load(); err != nil {
    log.Fatal("No .env file found")
  }

  spreadsheetId := os.Getenv("GOOGLE_SPREADS_SHEET_ID")
  readRange := os.Getenv("GOOGLE_SHEET_READ_RANGE")

  if spreadsheetId == "" {
    log.Fatal("GOOGLE_SPREADS_SHEET_ID не задан в переменных окружения")
  }

  if readRange == "" {
    log.Fatal("GOOGLE_SHEET_READ_RANGE не задан в переменных окружения")
  }

  log.Print("Config loaded")

	return &Config{
	  SpreadsSheetID: spreadsheetId,
    ReadRange: readRange,
	}
}