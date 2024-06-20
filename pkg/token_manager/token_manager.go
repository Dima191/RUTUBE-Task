package token_manager

import (
	"github.com/Dima191/RUTUBE-Task/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateAccessToken(employeeID uint32) (string, error)
	GenerateRefreshToken() (models.Session, error)
	Parse(accessToken string) (jwt.Claims, error)
}
