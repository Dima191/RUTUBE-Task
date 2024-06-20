package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/Dima191/RUTUBE-Task/pkg/hash"
	"github.com/google/uuid"
	"log/slog"
)

func (s *service) SignUp(ctx context.Context, employee models.SignUp) (accessToken string, refreshToken string, err error) {
	if err = s.validateErrorHandler(s.v.StructCtx(ctx, employee)); err != nil {
		return "", "", err
	}

	hashed, err := hash.Password(employee.Password)
	if err != nil {
		s.logger.Error("failed to hash password", slog.String("error", err.Error()))
		return "", "", srv.ErrInternal
	}

	employee.Password = hashed
	employee.ID = uuid.New().ID()

	if err = s.rep.SignUp(ctx, employee); err != nil {
		if errors.Is(err, rep.ErrEmployeeAlreadyExists) {
			return "", "", err
		}
		return "", "", srv.ErrInternal
	}

	accessToken, session, err := s.generateTokens(employee.ID)
	if err != nil {
		return "", "", err
	}

	err = s.createSession(ctx, employee.ID, session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
