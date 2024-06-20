package employeesrvimpl

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) updateSession(ctx context.Context, employeeID uint32, session models.Session) (err error) {
	if err = s.rep.UpdateSession(ctx, employeeID, session); err != nil {
		if errors.Is(err, rep.ErrNoSession) {
			return err
		}
		return srv.ErrInternal
	}
	return nil
}
