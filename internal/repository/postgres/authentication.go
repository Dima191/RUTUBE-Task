package postgresrep

import (
	"context"
	"errors"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *repository) Authentication(ctx context.Context, email string) (employeeID uint32, hashedPassword string, err error) {
	q := "select employee_id, hash_password from employee where email = $1"
	if err = r.pool.QueryRow(ctx, q, email).Scan(&employeeID, &hashedPassword); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("no employees by provided email", slog.String("email", email), slog.String("error", err.Error()))
			return 0, "", rep.ErrEmployeeNotFound
		}
		r.logger.Error("failed to find employee", slog.String("error", err.Error()))
		return 0, "", rep.ErrDatabaseQuery
	}

	return employeeID, hashedPassword, nil
}
