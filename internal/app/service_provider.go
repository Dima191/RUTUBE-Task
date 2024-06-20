package app

import (
	"context"
	geminiclient "github.com/Dima191/RUTUBE-Task/internal/clients/gemini"
	geminiclientimpl "github.com/Dima191/RUTUBE-Task/internal/clients/gemini/implmentation"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	postgresrep "github.com/Dima191/RUTUBE-Task/internal/repository/postgres"
	employeesrv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	employeesrvimpl "github.com/Dima191/RUTUBE-Task/internal/service/employee/implementation"
	msggenerator "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator"
	msggeneratorimpl "github.com/Dima191/RUTUBE-Task/internal/service/msg_generator/implementation"
	"github.com/Dima191/RUTUBE-Task/internal/service/notification"
	notificationimpl "github.com/Dima191/RUTUBE-Task/internal/service/notification/implementation"
	"github.com/Dima191/RUTUBE-Task/internal/workers/notifications"
	smtpmanager "github.com/Dima191/RUTUBE-Task/pkg/smtp_manager"
	smtpmanagerimpl "github.com/Dima191/RUTUBE-Task/pkg/smtp_manager/implementation"
	tm "github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	tokenmanagerimpl "github.com/Dima191/RUTUBE-Task/pkg/token_manager/implementation"
	"log/slog"
	"time"
)

type serviceProvider struct {
	repository rep.Repository

	employeeService     employeesrv.Service
	msgGeneratorService msggenerator.Service
	notificationService notification.Service

	geminiCl geminiclient.Client

	tokenManager tm.TokenManager

	smtpManager smtpmanager.Manager

	notificationWorker *notifications.Notification

	jwtSignedKey    []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	dbConnectionStr string

	geminiApiKey string

	smtpHost      string
	smtpPort      int
	senderEmail   string
	emailPassword string

	logger *slog.Logger
}

func (sp *serviceProvider) TokenManager() tm.TokenManager {
	if sp.tokenManager == nil {

		sp.tokenManager = tokenmanagerimpl.New(sp.jwtSignedKey, sp.accessTokenTTL, sp.refreshTokenTTL, sp.logger)
	}

	return sp.tokenManager
}

func (sp *serviceProvider) PostgresRep(ctx context.Context) (rep.Repository, error) {
	if sp.repository == nil {
		repository, err := postgresrep.New(ctx, sp.dbConnectionStr, sp.logger)
		if err != nil {
			return nil, err
		}
		sp.repository = repository
	}

	return sp.repository, nil
}

func (sp *serviceProvider) EmployeeService(ctx context.Context) (employeesrv.Service, error) {
	if sp.employeeService == nil {
		repository, err := sp.PostgresRep(ctx)
		if err != nil {
			return nil, err
		}
		sp.employeeService = employeesrvimpl.New(repository, sp.TokenManager(), sp.logger)
	}

	return sp.employeeService, nil
}

func (sp *serviceProvider) Gemini(ctx context.Context) (geminiclient.Client, error) {
	if sp.geminiCl == nil {
		geminiCl, err := geminiclientimpl.New(ctx, sp.geminiApiKey, sp.logger)
		if err != nil {
			return nil, err
		}
		sp.geminiCl = geminiCl
	}

	return sp.geminiCl, nil
}

func (sp *serviceProvider) SmtpManager() (smtpmanager.Manager, error) {
	if sp.smtpManager == nil {
		smtpManager, err := smtpmanagerimpl.New(sp.senderEmail, sp.emailPassword, sp.smtpHost, sp.smtpPort, sp.logger)
		if err != nil {
			return nil, err
		}

		sp.smtpManager = smtpManager
	}

	return sp.smtpManager, nil
}

func (sp *serviceProvider) MessageGeneratorService(ctx context.Context) (msggenerator.Service, error) {
	if sp.msgGeneratorService == nil {
		geminiCl, err := sp.Gemini(ctx)
		if err != nil {
			return nil, err
		}

		sp.msgGeneratorService = msggeneratorimpl.New(geminiCl, sp.logger)
	}

	return sp.msgGeneratorService, nil
}

func (sp *serviceProvider) NotificationService(ctx context.Context) (notification.Service, error) {
	if sp.notificationService == nil {
		msgGeneratorService, err := sp.MessageGeneratorService(ctx)
		if err != nil {
			return nil, err
		}

		smtpManager, err := sp.SmtpManager()
		if err != nil {
			return nil, err
		}

		sp.notificationService = notificationimpl.New(msgGeneratorService, smtpManager, sp.logger)
	}

	return sp.notificationService, nil
}

func (sp *serviceProvider) NotificationWorker(ctx context.Context) (*notifications.Notification, error) {
	if sp.notificationWorker == nil {
		notificationService, err := sp.NotificationService(ctx)
		if err != nil {
			return nil, err
		}

		employeeService, err := sp.EmployeeService(ctx)
		if err != nil {
			return nil, err
		}

		sp.notificationWorker = notifications.New(notificationService, employeeService, sp.logger)
	}

	return sp.notificationWorker, nil
}

func newServiceProvider(jwtSignedKey []byte, accessTokenTTL, refreshTokenTTL time.Duration, geminiApiKey string, dbConnectionStr string, smtpHost string, smtpPort int, senderEmail string, emailPassword string, logger *slog.Logger) *serviceProvider {
	sp := &serviceProvider{
		jwtSignedKey:    jwtSignedKey,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		dbConnectionStr: dbConnectionStr,
		geminiApiKey:    geminiApiKey,
		smtpHost:        smtpHost,
		smtpPort:        smtpPort,
		senderEmail:     senderEmail,
		emailPassword:   emailPassword,
		logger:          logger,
	}
	return sp
}
