// internal/service/sheet_mapper.go
package service

import (
	"fmt"
	"google_sheets_api/internal/domain"
)

func MapRowsToEvents(raw [][]interface{}) ([]domain.Event, error) {
	events := make([]domain.Event, 0, len(raw))

	for _, r := range raw {
		if len(r) < 2 {
			continue 
		}

		date := fmt.Sprint(r[0])
		name := fmt.Sprint(r[1])

		events = append(events, domain.Event{
			Date: date,
			Name: name,
		})
	}

	return events, nil
}