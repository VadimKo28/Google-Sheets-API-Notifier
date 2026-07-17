package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SpreadsSheetID  string
	ReadRange       string
	AppPassword     string
	GmailUser       string
	PostgresConnStr string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	spreadsheetId := os.Getenv("GOOGLE_SPREADS_SHEET_ID")
	readRange := os.Getenv("GOOGLE_SHEET_READ_RANGE")
	appPassword := os.Getenv("GMAIL_APP_PASSWORD")
	gmailUser := os.Getenv("GMAIL_USER")
	postgresConnStr := os.Getenv("DB_STRING")

	if postgresConnStr == "" {
		log.Fatal("DB_STRING не задан в переменных окружения")
	}

	if gmailUser == "" {
		log.Fatal("GMAIL_USER не задан в переменных окружения")
	}

	if appPassword == "" {
		log.Fatal("GMAIL_APP_PASSWORD не задан в переменных окружения")
	}

	if spreadsheetId == "" {
		log.Fatal("GOOGLE_SPREADS_SHEET_ID не задан в переменных окружения")
	}

	if readRange == "" {
		log.Fatal("GOOGLE_SHEET_READ_RANGE не задан в переменных окружения")
	}

	return &Config{
		SpreadsSheetID:  spreadsheetId,
		ReadRange:       readRange,
		AppPassword:     appPassword,
		GmailUser:       gmailUser,
		PostgresConnStr: postgresConnStr,
	}
}
