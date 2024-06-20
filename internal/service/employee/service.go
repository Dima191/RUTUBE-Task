package employeesrv

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
)

type Service interface {
	SignUp(ctx context.Context, employee models.SignUp) (accessToken string, refreshToken string, err error)
	SignIn(ctx context.Context, credentials models.SignIn) (accessToken string, refreshToken string, err error)
	UpdateTokens(ctx context.Context, expiredRefreshToken string) (accessToken string, refreshToken string, err error)
	LogOut(ctx context.Context, refreshToken string) error

	EmployeeByID(ctx context.Context, employeeID uint32) (models.Employee, error)
	Employees(ctx context.Context) ([]models.Employee, error)
	Subscribe(ctx context.Context, subscriberID, targetID uint32) error
	Unsubscribe(ctx context.Context, subscriberID, targetID uint32) error
	Subscriptions(ctx context.Context, subscriberID uint32) ([]models.Employee, error)

	TodayBirthdays(ctx context.Context) (notifications []models.Notify, err error)
}
