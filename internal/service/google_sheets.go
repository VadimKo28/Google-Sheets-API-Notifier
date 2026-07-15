package service

import (
	"context"
	"google_sheets_api/internal/domain"
	mapper "google_sheets_api/internal/lib/google"
	"google_sheets_api/internal/repository/event"
	"google_sheets_api/pkg/clients/google_sheets"
	"log/slog"

)

type GoogleSheetsService struct {
  client *google_sheets.GoogleSheetsClient
  logger *slog.Logger
  repository *event.EventRepository
}

func NewGoogleSheetsService(ctx context.Context, spreadsheetId string, readRange string, logger *slog.Logger, repository *event.EventRepository) (*GoogleSheetsService, error) {
  client, err := google_sheets.NewClient(ctx, spreadsheetId, readRange, logger)

  if err != nil {
    logger.Error("error create service", slog.Any("error", err))
    return nil, err
  }

  return &GoogleSheetsService{client: client, logger: logger, repository: repository}, nil
}

func (s *GoogleSheetsService) GetSheets(ctx context.Context) ([]domain.GoogleSheetElement, error){
  v, err := s.client.GetValues()

  if err != nil {
    s.logger.Error("error get values", slog.Any("error", err))
    return nil, err
  }

  rawData := v.Values

  mapData, err := mapper.MapRowsToEvents(rawData)

  return mapData, err
}

func (s *GoogleSheetsService) SaveEvents(ctx context.Context, events []domain.GoogleSheetElement) error {

  if err := s.repository.Save(ctx, events); err != nil {
    s.logger.Error("error save events", slog.Any("error", err))
    return err
  }

  return nil
}

func (s *GoogleSheetsService) SyncSheets(ctx context.Context) error{
  sheets, err := s.GetSheets(ctx)

  if err != nil {
    s.logger.Error("error sync sheets", slog.Any("error", err))
    return err
  }

  s.SaveEvents(ctx, sheets)

  return nil
}



func (s *GoogleSheetsService) SendMail() {
  panic("not implement")
}

