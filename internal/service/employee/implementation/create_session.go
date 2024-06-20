package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) createSession(ctx context.Context, employeeID uint32, session models.Session) error {
	if err := s.rep.CreateSession(ctx, employeeID, session); err != nil {
		return srv.ErrInternal
	}

	return nil
}
