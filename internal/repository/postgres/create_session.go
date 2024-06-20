package postgresrep

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) CreateSession(ctx context.Context, employeeID uint32, session models.Session) error {
	q := "insert into session(employee_id, refresh_token, expires_at) values ($1, $2, $3)"

	if _, err := r.pool.Exec(ctx, q, employeeID, session.Token, session.ExpiresAt); err != nil {
		r.logger.Error("failed to create session", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return nil
}
