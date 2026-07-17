package config

import (
	"log"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	SpreadsSheetID  string `env:"GOOGLE_SPREADS_SHEET_ID,notEmpty"`
	ReadRange       string `env:"GOOGLE_SHEET_READ_RANGE,notEmpty"` 
	AppPassword     string `env:"GMAIL_APP_PASSWORD,notEmpty"`
	GmailUser       string `env:"GMAIL_USER,notEmpty"`
	PostgresConnStr string `env:"DB_STRING,notEmpty"`
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := Config{} 

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
