package employeesrvimpl

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestLogOut(t *testing.T) {
	const refreshToken = "refresh token"

	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockedRep.EXPECT().LogOut(gomock.Any(), refreshToken).Return(nil)

	err := s.LogOut(ctx, refreshToken)

	assert.NoError(t, err)
}

func TestLogOutErr(t *testing.T) {
	const refreshToken = "refresh token"

	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		expectedErr error
		prepareFunc func()
	}{
		{
			name:        "INTERNAL ERROR",
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().LogOut(gomock.Any(), refreshToken).Return(rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			err := s.LogOut(ctx, refreshToken)
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
