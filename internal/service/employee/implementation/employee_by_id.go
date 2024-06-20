package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) EmployeeByID(ctx context.Context, employeeID uint32) (models.Employee, error) {
	employee, err := s.rep.EmployeeByID(ctx, employeeID)
	if err != nil {
		if errors.Is(err, rep.ErrEmployeeNotFound) {
			return models.Employee{}, err
		}
		return models.Employee{}, srv.ErrInternal
	}

	return employee, nil
}
