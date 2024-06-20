package postgresrep

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) Subscribe(ctx context.Context, subscriberID, targetID uint32) error {
	q := "insert into employee_subscription(subscriber_id, target_id) values($1,$2)"

	if _, err := r.pool.Exec(ctx, q, subscriberID, targetID); err != nil {
		r.logger.Error("failed to subscribe", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return nil
}
