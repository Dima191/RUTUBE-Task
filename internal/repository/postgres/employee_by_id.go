package postgresrep

import (
	"context"
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *repository) EmployeeByID(ctx context.Context, employeeID uint32) (models.Employee, error) {
	q := "select employee_id, full_name, birth_date, email from employee where employee_id = $1"

	employee := models.Employee{}
	var birthDate time.Time

	if err := r.pool.QueryRow(ctx, q, employeeID).Scan(&employee.ID, &employee.FullName, &birthDate, &employee.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Error("no employees", slog.String("error", err.Error()))
			return models.Employee{}, rep.ErrEmployeeNotFound
		}
		r.logger.Error("failed to find user", slog.String("error", err.Error()))
		return models.Employee{}, rep.ErrDatabaseQuery
	}
	employee.BirthDate = models.CustomDate{Time: birthDate}
	return employee, nil

}
