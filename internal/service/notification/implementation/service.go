package notificationimpl

import (
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	"github.com/Dima191/RUTUBE-Task/internal/service/notification"
	"github.com/Dima191/RUTUBE-Task/pkg/smtp_manager"
	"log/slog"
)

var (
	introduction = "Don't forget about your colleague' birth day\n"
)

type service struct {
	rep         rep.Repository
	smtpManager smtpmanager.Manager

	msgGenerator msggenerator.Service

	logger *slog.Logger
}

func New(msgGenerator msggenerator.Service, smtpManager smtpmanager.Manager, logger *slog.Logger) notification.Service {
	s := &service{
		smtpManager:  smtpManager,
		msgGenerator: msgGenerator,
		logger:       logger,
	}

	return s
}
