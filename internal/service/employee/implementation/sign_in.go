package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

func (s *service) SignIn(ctx context.Context, credentials models.SignIn) (accessToken string, refreshToken string, err error) {
	if err = s.validateErrorHandler(s.v.StructCtx(ctx, credentials)); err != nil {
		return "", "", err
	}

	employeeID, hashed, err := s.rep.Authentication(ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, rep.ErrEmployeeNotFound) {
			return "", "", srv.ErrInvalidLoginOrPassword
		}
		return "", "", srv.ErrInternal
	}

	passwordBytes := unsafe.Slice(unsafe.StringData(credentials.Password), len(credentials.Password))
	hashedPasswordPrt := unsafe.Slice(unsafe.StringData(hashed), len(hashed))
	if err = bcrypt.CompareHashAndPassword(hashedPasswordPrt, passwordBytes); err != nil {
		return "", "", srv.ErrInvalidLoginOrPassword
	}

	accessToken, session, err := s.generateTokens(employeeID)
	if err != nil {
		return "", "", err
	}

	err = s.updateSession(ctx, employeeID, session)
	switch {
	case err == nil:
		return accessToken, refreshToken, nil
	case errors.Is(err, rep.ErrNoSession):
	default:
		return "", "", srv.ErrInternal
	}

	err = s.createSession(ctx, employeeID, session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
