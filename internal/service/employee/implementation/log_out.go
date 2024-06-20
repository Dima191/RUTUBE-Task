package employeesrvimpl

import (
	"context"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) LogOut(ctx context.Context, refreshToken string) error {
	if err := s.rep.LogOut(ctx, refreshToken); err != nil {
		return srv.ErrInternal
	}

	return nil
}
