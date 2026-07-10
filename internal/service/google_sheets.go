package service

import (
	"context"
	"google_sheets_api/internal/client/google_sheets"
	"google_sheets_api/internal/domain"
)

type GoogleSheetsService struct {
  client *google_sheets.GoogleSheetsClient
}

func NewGoogleSheetsService(ctx context.Context, spreadsheetId string, readRange string) *GoogleSheetsService {
  client := google_sheets.NewGoogleSheetsClient(ctx, spreadsheetId, readRange)

  return &GoogleSheetsService{client: client}
}

func (s *GoogleSheetsService) GetAndMappingSheets(ctx context.Context) ([]domain.Event, error){
  v, _ := s.client.GetValues()

  return MapRowsToEvents(v.Values)
}



func (s *GoogleSheetsService) Notify() {
  panic("not implement")
}

