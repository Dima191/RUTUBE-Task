package employeesrvimpl

import (
	"context"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) Unsubscribe(ctx context.Context, subscriberID, targetID uint32) error {
	if err := s.rep.Unsubscribe(ctx, subscriberID, targetID); err != nil {
		return srv.ErrInternal
	}

	return nil
}
