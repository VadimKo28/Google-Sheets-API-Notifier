package event

import (
	"context"
	"fmt"
	"google_sheets_api/internal/domain"
	"log/slog"
	"time"

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
		INSERT INTO events (event_date, name, complete)
		VALUES ($1, $2, $3)
		ON CONFLICT (event_date) DO NOTHING`

		
	for i, row := range rows {
		if _, err := tx.Exec(ctx, query, row.Date, row.Name, row.Complete); err != nil {
      r.logger.Error("failed to insert row", slog.Any("error", err))
			return fmt.Errorf("failed to insert row %d (%s): %w", i, row.Name, err)
		}
	}

	return tx.Commit(ctx)
}

func (r *EventRepository) GetByDate(ctx context.Context, date time.Time) ([]domain.GoogleSheetElement, error) {
  dateStr := date.Format("02.01")

	rows, err := r.db.Query(ctx,
		`SELECT event_date, name, complete FROM events WHERE event_date = $1`,
		dateStr,
	)
	if err != nil {
    r.logger.Error("failed to query events", slog.Any("error", err))
		return nil, fmt.Errorf("failed to query events: %w", err)
	}
	defer rows.Close()

	var events []domain.GoogleSheetElement
	for rows.Next() {
		var e domain.GoogleSheetElement
		if err := rows.Scan(&e.Date, &e.Name, &e.Complete); err != nil {
      r.logger.Error("failed to scan event", slog.Any("error", err))
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}
		events = append(events, e)
	}
	if err := rows.Err(); err != nil {
    r.logger.Error("rows error", slog.Any("error", err))
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return events, nil
}
