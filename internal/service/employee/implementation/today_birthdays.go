package employeesrvimpl

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
)

func (s *service) TodayBirthdays(ctx context.Context) (notifications []models.Notify, err error) {
	notifications, err = s.rep.TodayBirthdays(ctx)
	if err != nil {
		return nil, srv.ErrInternal
	}

	return notifications, nil
}
