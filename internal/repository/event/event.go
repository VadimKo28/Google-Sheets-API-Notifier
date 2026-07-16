package event

import (
	"context"
	"fmt"
	"google_sheets_api/internal/domain"
	"log/slog"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepository struct {
    db *pgxpool.Pool 
	logger *slog.Logger
}

func NewEventRepository(db *pgxpool.Pool, logger *slog.Logger) *EventRepository {
    return &EventRepository{db: db, logger: logger}
}

func (r *EventRepository) Save(ctx context.Context, rows []domain.GoogleSheetElement) error {
	if len(rows) == 0 {
		return nil
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
	  r.logger.Error("failed to begin tx", slog.Any("error", err))
	  return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	const query = `
		INSERT INTO events (event_date, name, execute)
		VALUES ($1, $2, $3)
		ON CONFLICT (event_date) DO NOTHING`

		
	for i, row := range rows {
		if _, err := tx.Exec(ctx, query, row.Date, row.Name, row.Execute); err != nil {
      r.logger.Error("failed to insert row", slog.Any("error", err))
			return fmt.Errorf("failed to insert row %d (%s): %w", i, row.Name, err)
		}
	}

	return tx.Commit(ctx)
}

func (r *EventRepository) GetAll(ctx context.Context) ([]domain.GoogleSheetElement, error) {
  query := `
    SELECT event_date, name, execute
    FROM events
  `
  rows, err := r.db.Query(ctx, query)
  
  if err != nil {
    r.logger.Error("error get all events", slog.Any("error", err))
    return nil, err
  }
  defer rows.Close()

  events := make([]domain.GoogleSheetElement, 0)

  for rows.Next() {
    var event domain.GoogleSheetElement
    if err := rows.Scan(&event.Date, &event.Name, &event.Execute); err != nil {
      r.logger.Error("error scan row", slog.Any("error", err))
      return nil, err
    }
    events = append(events, event)
  }

  return events, nil
}