package event

import (
	"context"
	"fmt"
	"google_sheets_api/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepository struct {
    db *pgxpool.Pool 
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
    return &EventRepository{db: db}
}

func (r *EventRepository) Save(ctx context.Context, rows []domain.GoogleSheetElement) error {
	if len(rows) == 0 {
		return nil
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}
	for _, row := range rows {
		batch.Queue(
			`INSERT INTO events (event_date, name, execute) 
			 VALUES ($1, $2, $3)
			 ON CONFLICT (event_date) DO NOTHING`,
			row.Date, row.Name, row.Execute,
		)
	}

	br := tx.SendBatch(ctx, batch)
	for i := 0; i < len(rows); i++ {
		if _, err := br.Exec(); err != nil {
			br.Close()
			return fmt.Errorf("failed to insert row %d (%s): %w", i, rows[i].Name, err)
		}
	}
	if err := br.Close(); err != nil {
		return fmt.Errorf("failed to close batch: %w", err)
	}

	return tx.Commit(ctx)
}
