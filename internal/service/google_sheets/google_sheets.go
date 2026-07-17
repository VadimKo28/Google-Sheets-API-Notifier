package google_sheets

import (
	"context"
	"google_sheets_api/internal/domain"
	mapper "google_sheets_api/internal/lib/google"
	"google_sheets_api/internal/repository/event"
	"log/slog"

	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsService struct {
	sheetsClient SheetsClient
	logger       *slog.Logger
	repository   *event.EventRepository
}

type SheetsClient interface {
	GetValues() (*sheets.ValueRange, error)
}

func NewGoogleSheetsService(sheetsClient SheetsClient, logger *slog.Logger, repository *event.EventRepository) (*GoogleSheetsService, error) { // (client *SheetsClient, logger *slog.Logger, repository *event.EventRepository) (*GoogleSheetsService, error) {

	return &GoogleSheetsService{sheetsClient: sheetsClient, logger: logger, repository: repository}, nil
}

func (s *GoogleSheetsService) GetSheets(ctx context.Context) ([]domain.GoogleSheetElement, error) {
	v, err := s.sheetsClient.GetValues()

	if err != nil {
		s.logger.Error("error get values", slog.Any("error", err))
		return nil, err
	}

	rawData := v.Values

	mapData, err := mapper.MapRowsToEvents(rawData)

	if err != nil {
		s.logger.Error("error map rows to events", slog.Any("error", err))
		return nil, err
	}

	return mapData, err
}

func (s *GoogleSheetsService) SaveSheets(ctx context.Context, events []domain.GoogleSheetElement) error {

	if err := s.repository.Save(ctx, events); err != nil {
		s.logger.Error("error save events", slog.Any("error", err))
		return err
	}

	s.logger.Info("save events success")
	return nil
}

func (s *GoogleSheetsService) SyncSheets(ctx context.Context) error {
	sheets, err := s.GetSheets(ctx)

	if err != nil {
		s.logger.Error("error sync sheets", slog.Any("error", err))
		return err
	}

	err = s.SaveSheets(ctx, sheets)

	if err != nil {
		s.logger.Error("error sync sheets", slog.Any("error", err))
		return err
	}

	s.logger.Info("sync sheets success")
	return nil
}
