package mail

import (
	"log/slog"

	"gopkg.in/gomail.v2"
)

type MailClient struct {
  dialer *gomail.Dialer
  from string
  logger *slog.Logger
}

func NewMailClient(username, appPassword, from string, logger *slog.Logger) *MailClient {
	return &MailClient{
		dialer: gomail.NewDialer("smtp.gmail.com", 587, username, appPassword),
		from:   from,
		logger: logger,
	}
}

func (c *MailClient) Send(to string, subject, htmlBody string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", c.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	if err := c.dialer.DialAndSend(m); err != nil {
		c.logger.Error("failed to send email",
			slog.Any("error", err),
			slog.Any("to", to),
			slog.String("subject", subject))
		return err
	}

	return nil
}