package repository

import (
	"context"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	"time"
)

type Repository interface {
	SignUp(ctx context.Context, employee models.SignUp) error
	Authentication(ctx context.Context, email string) (employeeID uint32, hashedPassword string, err error)
	UpdateSession(ctx context.Context, employeeID uint32, session models.Session) error
	CreateSession(ctx context.Context, employeeID uint32, session models.Session) error
	RefreshTokenExpiration(ctx context.Context, refreshToken string) (employeeID uint32, expiredAt time.Time, err error)
	LogOut(ctx context.Context, refreshToken string) error

	Employees(ctx context.Context) ([]models.Employee, error)
	EmployeeByID(ctx context.Context, employeeID uint32) (models.Employee, error)
	Subscribe(ctx context.Context, subscriberID, targetID uint32) error
	Unsubscribe(ctx context.Context, subscriberID, targetID uint32) error
	Subscriptions(ctx context.Context, subscriberID uint32) ([]models.Employee, error)
	CheckSubscription(ctx context.Context, subscriberID, targetID uint32) error

	TodayBirthdays(ctx context.Context) (notifications []models.Notify, err error)

	CloseConnection()
}
