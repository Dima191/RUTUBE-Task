package postgresrep

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) Subscriptions(ctx context.Context, subscriberID uint32) ([]models.Employee, error) {
	employees := []models.Employee{}

	q := "select employee_id, target_full_name, target_email, target_birth_date from employee_subscriptions_by_name where subscriber_id = $1"

	rows, err := r.pool.Query(ctx, q, subscriberID)
	if err != nil {
		r.logger.Error("failed to get subscriptions", slog.String("error", err.Error()))
		return nil, rep.ErrDatabaseQuery
	}

	for rows.Next() {
		employee := models.Employee{}
		if err = rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.BirthDate); err != nil {
			r.logger.Error("failed to get subscriptions", slog.String("error", err.Error()))
			return nil, rep.ErrDatabaseQuery
		}

		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("failed to get subscriptions", slog.String("error", err.Error()))
		return nil, rep.ErrDatabaseQuery
	}

	return employees, nil
}
