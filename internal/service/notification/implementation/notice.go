package notificationimpl

import (
	"context"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/notification"
	"log/slog"
	"strings"
)

func (s *service) Notice(ctx context.Context, subscriberFullName, subscriberEmail string, celebrantFullName, celebrantEmail string) error {
	message := strings.Builder{}
	message.WriteString(introduction)
	message.WriteString("Celebrant email: ")
	message.WriteString(celebrantEmail)

	congratulation, err := s.msgGenerator.Generate(ctx, subscriberFullName, celebrantFullName)
	switch {
	case err != nil:
		s.logger.Warn("failed to generate congratulatory message", slog.String("error", err.Error()))
	default:
		message.WriteString("\nSuggested congratulation:\n")
		message.WriteString(congratulation)
	}

	err = s.smtpManager.SendMail(subscriberEmail, message.String())
	if err != nil {
		return srv.ErrInternal
	}

	return nil
}
