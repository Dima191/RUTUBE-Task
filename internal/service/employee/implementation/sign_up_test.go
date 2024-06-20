package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSignUp(t *testing.T) {
	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cases := []struct {
		name        string
		credentials models.SignUp
		prepareFunc func()
	}{
		{
			name: "OK",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "valid_email@gmail.com",
				},
				Password: "valid_password1234",
			},
			prepareFunc: func() {
				mockedRep.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
				mockedRep.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(gomock.Any()).Return("access token", nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, _, err := s.SignUp(ctx, c.credentials)
			assert.Nil(t, err)
		})
	}
}

func TestSignUpErr(t *testing.T) {
	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name                 string
		credentials          models.SignUp
		expectedErr          error
		expectedErrInterface interface{}
		prepareFunc          func()
	}{
		{
			name: "INVALID EMAIL",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "invalid_email@gmail_com",
				},
				Password: "valid_password1234",
			},
			expectedErrInterface: validator.ValidationErrors{},
		},
		{
			name: "INVALID PASSWORD",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "valid_email@gmail.com",
				},
				Password: "1234567",
			},
			expectedErrInterface: validator.ValidationErrors{},
		},
		{
			name: "EMPLOYEE ALREADY EXISTS",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "valid_email@gmail.com",
				},
				Password: "valid_password1234",
			},
			expectedErr: rep.ErrEmployeeAlreadyExists,
			prepareFunc: func() {
				mockedRep.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(rep.ErrEmployeeAlreadyExists)
			},
		},
		{
			name: "INTERNAL ERROR",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "valid_email@gmail.com",
				},
				Password: "valid_password1234",
			},
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(rep.ErrDatabaseQuery)
			},
		},
		{
			name: "CREATE SESSION ERROR",
			credentials: models.SignUp{
				Employee: models.Employee{
					Email: "valid_email@gmail.com",
				},
				Password: "valid_password1234",
			},
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
				mockedRep.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(rep.ErrDatabaseQuery)
				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(gomock.Any()).Return("access token", nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.prepareFunc != nil {
				c.prepareFunc()
			}
			_, _, err := s.SignUp(context.Background(), c.credentials)
			if errors.As(err, &c.expectedErrInterface) {
				return
			}
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
