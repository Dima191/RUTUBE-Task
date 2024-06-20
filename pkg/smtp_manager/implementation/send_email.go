package smtpmanagerimpl

import (
	"gopkg.in/gomail.v2"
	"log/slog"
)

func (m *smtpManager) SendMail(to string, msg string) error {
	message := gomail.NewMessage()

	m.logger.Info("sending email")

	message.SetHeader("From", m.email)
	message.SetHeader("To", to)
	message.SetHeader("Subject", EmailTheme)
	message.SetBody("text/plain", msg)

	if err := m.dialer.DialAndSend(message); err != nil {
		m.logger.Error("failed to send message", slog.String("error", err.Error()))
		return err
	}
	return nil
}
