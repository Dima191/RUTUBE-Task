package smtpmanagerimpl

import (
	"crypto/tls"
	"github.com/Dima191/RUTUBE-Task/pkg/smtp_manager"
	"gopkg.in/gomail.v2"
	"log/slog"
)

const EmailTheme = "Colleague's birthday"

type smtpManager struct {
	dialer *gomail.Dialer

	email string
	host  string
	port  int

	logger *slog.Logger
}

func New(email string, password string, host string, port int, logger *slog.Logger) (smtpmanager.Manager, error) {
	dialer := gomail.NewDialer(host, port, email, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &smtpManager{
		dialer: dialer,
		email:  email,
		host:   host,
		port:   port,
		logger: logger,
	}, nil

}
