package google_sheets

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetsClient struct {
  googleClient *sheets.Service
  spreadsheetId string
  readRange string
}

func NewGoogleSheetsClient(ctx context.Context, spreadsheetId string, readRange string) *GoogleSheetsClient {
  srv, _ := sheets.NewService(ctx, option.WithAuthCredentialsFile(option.ServiceAccount, "my-sheets-integration-501715-8e3105270262.json"))

  return &GoogleSheetsClient{googleClient: srv, spreadsheetId: spreadsheetId, readRange: readRange}
}

func (srv *GoogleSheetsClient) GetValues() (*sheets.ValueRange, error) {
  resp, _ := srv.googleClient.Spreadsheets.Values.Get(srv.spreadsheetId, srv.readRange).Do()

	return resp, nil
}