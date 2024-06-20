package postgresrep

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) Unsubscribe(ctx context.Context, subscriberID, targetID uint32) error {
	q := "delete from employee_subscription where target_id = $1 and subscriber_id = $2"

	if _, err := r.pool.Exec(ctx, q, targetID, subscriberID); err != nil {
		r.logger.Error("failed to unsubscribe", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return nil
}
