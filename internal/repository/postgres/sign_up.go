package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"time"
)

func (r *repository) SignUp(ctx context.Context, employee models.SignUp) error {
	q := "insert into employee(employee_id, full_name, birth_date, email, hash_password) values ($1, $2, $3, $4, $5)"

	_, err := r.pool.Exec(ctx, q, employee.ID, employee.FullName, employee.BirthDate.Format(time.DateOnly), employee.Email, employee.Password)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return rep.ErrEmployeeAlreadyExists
			}
		}
		r.logger.Error("failed to register employee", slog.String("error", err.Error()))
		return rep.ErrDatabaseQuery
	}

	return nil
}
