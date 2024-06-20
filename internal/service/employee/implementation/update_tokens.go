package employeesrvimpl

import (
	"context"
	"errors"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"reflect"
	"time"
)

func (s *service) UpdateTokens(ctx context.Context, oldRefreshToken string) (accessToken string, refreshToken string, err error) {
	employeeID, expAt, err := s.rep.RefreshTokenExpiration(ctx, oldRefreshToken)
	if err != nil {
		if errors.Is(err, rep.ErrNoSession) {
			return "", "", err
		}

		return "", "", srv.ErrInternal
	}

	if expAt.Compare(time.Now()) <= 0 || reflect.DeepEqual(expAt, time.Time{}) {
		return "", "", srv.ErrTokenExpired
	}

	accessToken, session, err := s.generateTokens(employeeID)
	if err != nil {
		return "", "", err
	}

	err = s.updateSession(ctx, employeeID, session)
	if err != nil {
		return "", "", err
	}

	return accessToken, session.Token, nil
}
