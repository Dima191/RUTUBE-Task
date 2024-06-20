package employeesrvimpl

import (
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	tm "github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"github.com/go-playground/validator"
	"log/slog"
)

type service struct {
	rep          rep.Repository
	tokenManager tm.TokenManager
	v            *validator.Validate

	logger *slog.Logger
}

func New(rep rep.Repository, tokenManager tm.TokenManager, logger *slog.Logger) srv.Service {
	s := &service{
		rep:          rep,
		tokenManager: tokenManager,
		v:            validator.New(),
		logger:       logger,
	}

	return s
}
