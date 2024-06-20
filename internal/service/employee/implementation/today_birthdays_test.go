package employeesrvimpl

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTodayBirthdays(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockedRep.EXPECT().TodayBirthdays(gomock.Any()).Return(nil, nil)

	_, err := s.TodayBirthdays(ctx)
	assert.NoError(t, err)
}

func TestTodayBirthdaysErr(t *testing.T) {
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
				mockedRep.EXPECT().TodayBirthdays(gomock.Any()).Return(nil, rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, err := s.TodayBirthdays(ctx)
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
