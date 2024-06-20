package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) Subscriptions(ctx context.Context, subscriberID uint32) ([]models.Employee, error) {
	employees, err := s.rep.Subscriptions(ctx, subscriberID)
	if err != nil {
		return nil, srv.ErrInternal
	}

	return employees, nil
}
