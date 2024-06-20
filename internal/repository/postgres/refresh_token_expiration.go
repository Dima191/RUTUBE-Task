package postgresrep

import (
	"context"
	"errors"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *repository) RefreshTokenExpiration(ctx context.Context, refreshToken string) (employeeID uint32, expiredAt time.Time, err error) {
	q := "select employee_id, expires_at from session where refresh_token=$1"

	if err = r.pool.QueryRow(ctx, q, refreshToken).Scan(&employeeID, &expiredAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("no session by provided refresh token", slog.String("error", err.Error()))
			return 0, time.Time{}, rep.ErrNoSession
		}
		r.logger.Error("failed to get session", slog.String("error", err.Error()))
		return 0, time.Time{}, rep.ErrDatabaseQuery
	}

	return employeeID, expiredAt, nil
}
