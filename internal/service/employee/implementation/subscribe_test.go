package employeesrvimpl

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestSubscribe(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name         string
		subscriberID uint32
		targetID     uint32
	}{
		{name: "OK", subscriberID: 1, targetID: 2},
	}

	ctx := context.Background()
	mockedRep.EXPECT().Subscribe(gomock.Any(), uint32(1), uint32(2)).Return(nil)
	mockedRep.EXPECT().CheckSubscription(gomock.Any(), uint32(1), uint32(2)).Return(nil)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := s.Subscribe(ctx, c.subscriberID, c.targetID)
			assert.NoError(t, err)
		})
	}
}

func TestSubscribeErr(t *testing.T) {
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
			name:         "SELF SUBSCRIPTION",
			subscriberID: 1,
			targetID:     1,
			expectedErr:  srv.ErrSelfSubscription,
		},
		{
			name:         "ALREADY SUBSCRIBED",
			subscriberID: 1,
			targetID:     3,
			expectedErr:  rep.ErrAlreadySubscribed,
			prepareFunc: func() {
				mockedRep.EXPECT().CheckSubscription(gomock.Any(), uint32(1), uint32(3)).Return(rep.ErrAlreadySubscribed)
			},
		},
		{
			name:         "INTERNAL ERROR",
			subscriberID: 1,
			targetID:     3,
			expectedErr:  srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().CheckSubscription(gomock.Any(), uint32(1), uint32(3)).Return(rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.prepareFunc != nil {
				c.prepareFunc()
			}
			err := s.Subscribe(ctx, c.subscriberID, c.targetID)
			assert.Equal(t, c.expectedErr, err)
		})
	}
}
