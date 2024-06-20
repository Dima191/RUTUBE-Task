package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) Employees(ctx context.Context) ([]models.Employee, error) {
	employees, err := s.rep.Employees(ctx)
	if err != nil {
		return nil, srv.ErrInternal
	}

	return employees, nil
}
