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

func TestEmployeeByID(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		employeeID  uint32
		prepareFunc func()
	}{
		{
			name:       "OK",
			employeeID: 1,
			prepareFunc: func() {
				mockedRep.EXPECT().EmployeeByID(gomock.Any(), uint32(1)).Return(models.Employee{}, nil)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, err := s.EmployeeByID(ctx, c.employeeID)
			assert.NoError(t, err)
		})
	}
}

func TestEmployeeByIDErr(t *testing.T) {
	s, mockedRep, _, ctrl := testService(t)
	defer ctrl.Finish()

	cases := []struct {
		name        string
		employeeID  uint32
		expectedErr error
		prepareFunc func()
	}{
		{
			name:        "NO EMPLOYEE FOUND",
			employeeID:  1,
			expectedErr: rep.ErrEmployeeNotFound,
			prepareFunc: func() {
				mockedRep.EXPECT().EmployeeByID(gomock.Any(), uint32(1)).Return(models.Employee{}, rep.ErrEmployeeNotFound)
			},
		},
		{
			name:        "INTERNAL ERROR",
			employeeID:  1,
			expectedErr: srv.ErrInternal,
			prepareFunc: func() {
				mockedRep.EXPECT().EmployeeByID(gomock.Any(), uint32(1)).Return(models.Employee{}, rep.ErrDatabaseQuery)
			},
		},
	}

	ctx := context.Background()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.prepareFunc()
			_, err := s.EmployeeByID(ctx, c.employeeID)
			assert.ErrorIs(t, err, c.expectedErr)
		})
	}
}
