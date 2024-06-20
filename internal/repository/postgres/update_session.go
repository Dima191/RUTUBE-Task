package postgresrep

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) UpdateSession(ctx context.Context, employeeID uint32, session models.Session) error {
	q := "update session set refresh_token = $1, expires_at = $2 where employee_id = $3"

	tag, err := r.pool.Exec(ctx, q, session.Token, session.ExpiresAt, employeeID)
	if err != nil {
		r.logger.Error("failed to update session", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	if tag.RowsAffected() == 0 {
		r.logger.Error("employee has no session", slog.Int("employee", int(employeeID)))
		return rep.ErrNoSession
	}

	return nil
}
