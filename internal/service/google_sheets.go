package service

import (
	"context"
	"google_sheets_api/internal/domain"
	"google_sheets_api/pkg/clients/google"
	"log/slog"
  mapper "google_sheets_api/internal/lib/google"

)

type GoogleSheetsService struct {
  client *google.GoogleSheetsClient
  logger *slog.Logger
}

func NewGoogleSheetsService(ctx context.Context, spreadsheetId string, readRange string, logger *slog.Logger) (*GoogleSheetsService, error) {
  client, err := google.NewGoogleSheetsClient(ctx, spreadsheetId, readRange, logger)

  if err != nil {
    logger.Error("error create service", slog.Any("error", err))
    return nil, err
  }

  return &GoogleSheetsService{client: client, logger: logger}, nil
}

func (s *GoogleSheetsService) GetAndMappingSheets(ctx context.Context) ([]domain.Event, error){
  v, err := s.client.GetValues()

  if err != nil {
    s.logger.Error("error get values", slog.Any("error", err))
    return nil, err
  }

  return mapper.MapRowsToEvents(v.Values)
}



func (s *GoogleSheetsService) Notify() {
  panic("not implement")
}

