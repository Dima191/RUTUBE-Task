package employeesrvimpl

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestUnsubscribe(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	subscriberID, targetID := uint32(1), uint32(2)

	mockedRep.EXPECT().Unsubscribe(gomock.Any(), subscriberID, targetID).Return(nil)

	err := s.Unsubscribe(ctx, subscriberID, targetID)
	assert.NoError(t, err)
}

func TestUnsubscribeErr(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name         string
		subscriberID uint32
		targetID     uint32
		expectedErr  error
		prepareFunc  func()
	}{
		{
			name:         "INTERNAL ERROR",
			subscriberID: 1,
			targetID:     2,
			expectedErr:  srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().Unsubscribe(gomock.Any(), uint32(1), uint32(2)).Return(rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			err := s.Unsubscribe(ctx, c.subscriberID, c.targetID)
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
