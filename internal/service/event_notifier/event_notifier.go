package event_notifier

import (
	"context"
	"google_sheets_api/internal/config"
	"google_sheets_api/internal/domain"
	"log/slog"
	"time"
)

type EventNotifierService struct {
  logger *slog.Logger
  repository EventReader
  mailer Mailer
  config config.Config
}

type EventReader interface {
	GetByDate(ctx context.Context, date time.Time) ([]domain.GoogleSheetElement, error)
}

type Mailer interface {
	Send(to string, subject string, body string) error
}

func NewEventNotifierService(config config.Config, logger *slog.Logger, repository EventReader, mailer Mailer) *EventNotifierService {
  return &EventNotifierService{config: config, logger: logger, repository: repository, mailer: mailer}
}

func (s *EventNotifierService) CheckEventsToday(ctx context.Context) error {
  today := time.Now()

  events, err := s.repository.GetByDate(ctx, today)
  if err != nil {
	  s.logger.Error("error get events for today", slog.Any("error", err))
	  return err
  }

  if len(events) == 0 {
	  s.logger.Info("no events today")
	  return nil
	}

	for _, e := range events {
		if err := s.mailer.Send(s.config.GmailUser, "Событие сегодня", e.Name); err != nil {
			s.logger.Error("error send mail", slog.Any("error", err), slog.String("event", e.Name))
			return err
		}
	}

	s.logger.Info("events notified", slog.Int("count", len(events)))
	return nil
}
