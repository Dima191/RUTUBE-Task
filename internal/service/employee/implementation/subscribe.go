package employeesrvimpl

import (
	"context"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) Subscribe(ctx context.Context, subscriberID, targetID uint32) error {
	if subscriberID == targetID {
		return srv.ErrSelfSubscription
	}

	if err := s.checkSubscription(ctx, subscriberID, targetID); err != nil {
		return err
	}

	if err := s.rep.Subscribe(ctx, subscriberID, targetID); err != nil {
		return srv.ErrInternal
	}

	return nil
}
