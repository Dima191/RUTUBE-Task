package postgresrep

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) LogOut(ctx context.Context, refreshToken string) error {
	q := "delete from session where refresh_token = $1"
	if _, err := r.pool.Exec(ctx, q, refreshToken); err != nil {
		r.logger.Error("failed to log out", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return nil
}
