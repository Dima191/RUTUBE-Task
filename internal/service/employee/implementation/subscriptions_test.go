package employeesrvimpl

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSubscriptions(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockedRep.EXPECT().Subscriptions(gomock.Any(), uint32(1)).Return(nil, nil)

	_, err := s.Subscriptions(ctx, uint32(1))
	assert.NoError(t, err)
}

func TestSubscriptionsErr(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		employeeID  uint32
		expectedErr error
		prepareFunc func()
	}{
		{
			name:        "INTERNAL ERROR",
			employeeID:  uint32(1),
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().Subscriptions(gomock.Any(), uint32(1)).Return(nil, rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, err := s.Subscriptions(ctx, c.employeeID)
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
