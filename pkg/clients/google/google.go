package google

import (
	"context"
	"log/slog"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsClient struct {
  googleClient *sheets.Service
  spreadsheetId string
  readRange string
  logger *slog.Logger
}

func NewGoogleSheetsClient(ctx context.Context, spreadsheetId string, readRange string, logger *slog.Logger) (*GoogleSheetsClient, error) {
  srv, err := sheets.NewService(ctx, option.WithAuthCredentialsFile(option.ServiceAccount, "my-sheets-integration-501715-8e3105270262.json"))

  if err != nil {
    logger.Error("error create service", slog.Any("error", err))
    return nil, err
  }

  return &GoogleSheetsClient{googleClient: srv, spreadsheetId: spreadsheetId, readRange: readRange, logger: logger}, nil
}

func (srv *GoogleSheetsClient) GetValues() (*sheets.ValueRange, error) {
  resp, err := srv.googleClient.Spreadsheets.Values.Get(srv.spreadsheetId, srv.readRange).Do()
  
  if err != nil {
    srv.logger.Error("error get values", slog.Any("error", err))
    return nil, err
  }

	return resp, nil
}