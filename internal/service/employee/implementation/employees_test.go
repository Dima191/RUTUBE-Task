package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestEmployees(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockedRep.EXPECT().Employees(gomock.Any()).Return([]models.Employee{}, nil)

	_, err := s.Employees(ctx)
	assert.NoError(t, err)
}

func TestEmployeesErr(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name          string
		expectedError error
		prepareFunc   func()
	}{
		{
			name:          "INTERNAL ERROR",
			expectedError: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().Employees(gomock.Any()).Return(nil, rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, err := s.Employees(ctx)
			assert.ErrorIs(t, err, c.expectedError)
		})
	}
}
