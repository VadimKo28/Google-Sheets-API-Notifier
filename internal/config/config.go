package config

import (
	"log"
	"os"
)

type Config struct {
	SpreadsSheetID string
  ReadRange string
}

func NewConfig() *Config {
  spreadsheetId := os.Getenv("GOOGLE_SPREADS_SHEET_ID")
  readRange := os.Getenv("GOOGLE_SHEET_READ_RANGE")

  if spreadsheetId == "" {
    log.Fatal("GOOGLE_SPREADS_SHEET_ID не задан в переменных окружения")
  }

  if readRange == "" {
    log.Fatal("GOOGLE_SHEET_READ_RANGE не задан в переменных окружения")
  }

	return &Config{
	  SpreadsSheetID: spreadsheetId,
    ReadRange: readRange,
	}
}