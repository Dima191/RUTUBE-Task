package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/Dima191/RUTUBE-Task/pkg/hash"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSignIn(t *testing.T) {
	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		credentials models.SignIn
		prepareFunc func()
	}{
		{
			name: "UPDATE SESSION",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "valid_password1243",
			},
			prepareFunc: func() {
				hashedPassword, err := hash.Password("valid_password1243")

				require.NoError(t, err)
				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(1), hashedPassword, nil)
				mockedRep.EXPECT().UpdateSession(gomock.Any(), uint32(1), gomock.Any()).Return(nil)

				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh_token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(uint32(1)).Return("access_token", nil)

			},
		},
		{
			name: "CREATE SESSION",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "valid_password1243",
			},
			prepareFunc: func() {
				hashedPassword, err := hash.Password("valid_password1243")

				require.NoError(t, err)
				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(1), hashedPassword, nil)

				//NO SESSION TO UPDATE
				mockedRep.EXPECT().UpdateSession(gomock.Any(), uint32(1), gomock.Any()).Return(rep.ErrNoSession)

				//CREATE NEW SESSION
				mockedRep.EXPECT().CreateSession(gomock.Any(), uint32(1), gomock.Any()).Return(nil)

				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh_token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(uint32(1)).Return("access_token", nil)

			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.prepareFunc != nil {
				c.prepareFunc()
			}
			_, _, err := s.SignIn(ctx, c.credentials)
			assert.NoError(t, err)
		})
	}
}

func TestSignInErr(t *testing.T) {
	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name           string
		credentials    models.SignIn
		expectedErr    error
		errorInterface interface{}
		prepareFunc    func()
	}{
		{
			name: "UPDATE SESSION UNEXPECTED ERROR",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "valid_password1243",
			},
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				hashedPassword, err := hash.Password("valid_password1243")

				require.NoError(t, err)
				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(1), hashedPassword, nil)
				mockedRep.EXPECT().UpdateSession(gomock.Any(), uint32(1), gomock.Any()).Return(rep.ErrDatabaseQuery)

				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh_token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(uint32(1)).Return("access_token", nil)
			},
		},
		{
			name: "CREATE SESSION UNEXPECTED ERROR",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "valid_password1243",
			},
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				hashedPassword, err := hash.Password("valid_password1243")

				require.NoError(t, err)
				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(1), hashedPassword, nil)

				//NO SESSION TO UPDATE
				mockedRep.EXPECT().UpdateSession(gomock.Any(), uint32(1), gomock.Any()).Return(rep.ErrNoSession)

				//CREATE NEW SESSION
				mockedRep.EXPECT().CreateSession(gomock.Any(), uint32(1), gomock.Any()).Return(rep.ErrDatabaseQuery)

				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh_token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(uint32(1)).Return("access_token", nil)
			},
		},
		{
			name: "NO USER WITH THIS EMAIL",
			credentials: models.SignIn{
				Email:    "no_user@gmail.com",
				Password: "valid_password1243",
			},
			expectedErr: srv.ErrInvalidLoginOrPassword,
			prepareFunc: func() {
				mockedRep.EXPECT().Authentication(gomock.Any(), "no_user@gmail.com").Return(uint32(0), "", rep.ErrEmployeeNotFound)
			},
		},
		{
			name: "WRONG PASSWORD",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "wrong_password1243",
			},
			expectedErr: srv.ErrInvalidLoginOrPassword,
			prepareFunc: func() {
				hashedPassword, err := hash.Password("correct_password1243")
				require.NoError(t, err)

				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(1), hashedPassword, nil)
			},
		},
		{
			name: "AUTHENTICATION UNEXPECTED ERROR",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "valid_password1243",
			},
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().Authentication(gomock.Any(), "valid_email@gmail.com").Return(uint32(0), "", nil)
			},
		},
		{
			name: "INVALID EMAIL",
			credentials: models.SignIn{
				Email:    "invalid_email@gmail_com",
				Password: "valid_password1243",
			},
			errorInterface: validator.ValidationErrors{},
		},
		{
			name: "INVALID PASSWORD",
			credentials: models.SignIn{
				Email:    "valid_email@gmail.com",
				Password: "1234567",
			},
			errorInterface: validator.ValidationErrors{},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.prepareFunc != nil {
				c.prepareFunc()
			}
			_, _, err := s.SignIn(ctx, c.credentials)
			if errors.As(err, &c.errorInterface) {
				return
			}

			assert.ErrorIs(t, c.expectedErr, err)
		})
	}
}
