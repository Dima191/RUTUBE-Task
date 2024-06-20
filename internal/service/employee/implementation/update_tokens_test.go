package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestUpdateTokens(t *testing.T) {
	const refreshToken = "refresh token"

	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	expAt := time.Now().Add(5 * time.Minute)
	mockedRep.EXPECT().RefreshTokenExpiration(gomock.Any(), refreshToken).Return(uint32(1), expAt, nil)
	mockedRep.EXPECT().UpdateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh token"}, nil)
	mockedTokenManager.EXPECT().GenerateAccessToken(gomock.Any()).Return("access token", nil)

	_, _, err := s.UpdateTokens(ctx, refreshToken)
	assert.NoError(t, err)
}

func TestUpdateTokensErr(t *testing.T) {
	const refreshToken = "refresh token"

	s, mockedRep, mockedTokenManager, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name          string
		expectedError error
		prepareFunc   func()
	}{
		{
			name:          "NO SESSION TO UPDATE",
			expectedError: rep.ErrNoSession,
			prepareFunc: func() {
				mockedRep.EXPECT().RefreshTokenExpiration(gomock.Any(), refreshToken).Return(uint32(0), time.Time{}, rep.ErrNoSession)
			},
		},
		{
			name:          "RECEIVING REFRESH TOKEN EXPIRY DATE INTERNAL ERROR",
			expectedError: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().RefreshTokenExpiration(gomock.Any(), refreshToken).Return(uint32(0), time.Time{}, rep.ErrDatabaseQuery)
			},
		},
		{
			name:          "TOKEN EXPIRED",
			expectedError: srv.ErrTokenExpired,
			prepareFunc: func() {
				expAt := time.Date(time.Now().Year()-1, time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
				mockedRep.EXPECT().RefreshTokenExpiration(gomock.Any(), refreshToken).Return(uint32(1), expAt, nil)
			},
		},
		{
			name:          "UPDATE SESSION INTERNAL ERROR",
			expectedError: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().RefreshTokenExpiration(gomock.Any(), refreshToken).Return(uint32(1), time.Now().Add(5*time.Minute), nil)
				mockedRep.EXPECT().UpdateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(rep.ErrDatabaseQuery)
				mockedTokenManager.EXPECT().GenerateRefreshToken().Return(models.Session{Token: "refresh token"}, nil)
				mockedTokenManager.EXPECT().GenerateAccessToken(gomock.Any()).Return("access token", nil)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.prepareFunc != nil {
				c.prepareFunc()
			}

			_, _, err := s.UpdateTokens(ctx, refreshToken)
			assert.ErrorIs(t, err, c.expectedError)
		})
	}
}
