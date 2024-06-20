package postgresrep

import (
	"context"
	"errors"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *repository) CheckSubscription(ctx context.Context, subscriberID, targetID uint32) error {
	q := "select 1 from employee_subscription where target_id = $1 and subscriber_id = $2"

	flag := 0

	if err := r.pool.QueryRow(ctx, q, targetID, subscriberID).Scan(&flag); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}

		r.logger.Error("failed to check subscription", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return rep.ErrAlreadySubscribed
}
