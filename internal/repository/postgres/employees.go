package postgresrep

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
	"time"
	"unsafe"
)

func (r *repository) Employees(ctx context.Context) ([]models.Employee, error) {
	employees := make([]models.Employee, 0)

	q := "select employee_id, full_name, birth_date, email from employee"

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		r.logger.Error("failed to get employees", slog.String("error", err.Error()))
		return nil, rep.ErrDatabaseQuery
	}

	for rows.Next() {
		employee := models.Employee{}
		var birthDate time.Time

		if err = rows.Scan(&employee.ID, &employee.FullName, &employee.BirthDate, &employee.Email); err != nil {
			r.logger.Error("failed to scan employees", slog.String("error", err.Error()))
			return nil, rep.ErrDatabaseQuery
		}

		dataPtr := uintptr(unsafe.Pointer(&birthDate))
		employee.BirthDate = *(*models.CustomDate)(unsafe.Pointer(dataPtr))
		employees = append(employees, employee)
	}

	if rows.Err() != nil {
		r.logger.Error("failed to get employees", slog.String("error", rows.Err().Error()))
		return nil, rep.ErrDatabaseQuery
	}

	return employees, nil
}
