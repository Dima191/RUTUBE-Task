package postgresrep

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"log/slog"
)

func (r *repository) TodayBirthdays(ctx context.Context) (notifications []models.Notify, err error) {
	q := `select target_full_name,
       			 target_email,
       			 subscriber_full_name,
       			 subscriber_email 
		  from employee_subscriptions_by_name 
		  where 
		      EXTRACT (month from target_birth_date) = EXTRACT(MONTH FROM CURRENT_DATE) 
		      and EXTRACT (DAY from target_birth_date) = EXTRACT(DAY FROM CURRENT_DATE)`

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		r.logger.Error("failed to check employees' birth days", slog.String("error", err.Error()))
		return nil, rep.ErrDatabaseQuery
	}

	for rows.Next() {
		notice := models.Notify{}

		if err = rows.Scan(&notice.Celebrator.FullName, &notice.Celebrator.Email, &notice.Subscriber.FullName, &notice.Subscriber.Email); err != nil {
			r.logger.Error("failed to check employees' birth days", slog.String("error", err.Error()))
			return nil, rep.ErrDatabaseQuery
		}

		notifications = append(notifications, notice)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("failed to check employees' birth days", slog.String("error", err.Error()))
		return nil, rep.ErrDatabaseQuery
	}

	return notifications, nil

}
