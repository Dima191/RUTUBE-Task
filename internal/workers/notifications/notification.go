package notifications

import (
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/Dima191/RUTUBE-Task/internal/service/notification"
	"log/slog"
)

type Notification struct {
	srvNotification notification.Service
	srvEmployee     srv.Service

	logger *slog.Logger
}

func New(srvNotification notification.Service, srvEmployee srv.Service, logger *slog.Logger) *Notification {
	return &Notification{
		srvNotification: srvNotification,
		srvEmployee:     srvEmployee,
		logger:          logger,
	}
}
