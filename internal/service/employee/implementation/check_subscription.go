package employeesrvimpl

import (
	"context"
	"errors"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) checkSubscription(ctx context.Context, subscriberID, targetID uint32) error {
	if err := s.rep.CheckSubscription(ctx, subscriberID, targetID); err != nil {
		if errors.Is(err, rep.ErrAlreadySubscribed) {
			return err
		}
		return srv.ErrInternal
	}
	return nil
}
