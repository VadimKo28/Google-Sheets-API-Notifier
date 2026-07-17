// internal/service/sheet_mapper.go
package google

import (
	"fmt"
	"google_sheets_api/internal/domain"
)

func MapRowsToEvents(raw [][]interface{}) ([]domain.GoogleSheetElement, error) {
	events := make([]domain.GoogleSheetElement, 0, len(raw))

	for _, r := range raw {
		if len(r) < 2 {
			continue
		}

		date := fmt.Sprint(r[0])
		name := fmt.Sprint(r[1])

		events = append(events, domain.GoogleSheetElement{
			Date:     date,
			Name:     name,
			Complete: false,
		})
	}

	return events, nil
}
