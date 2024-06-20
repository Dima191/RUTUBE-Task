package employeesrvimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/models"
)

func (s *service) generateTokens(employeeID uint32) (accessToken string, session models.Session, err error) {
	accessToken, err = s.tokenManager.GenerateAccessToken(employeeID)
	if err != nil {
		return "", models.Session{}, err
	}

	session, err = s.tokenManager.GenerateRefreshToken()
	if err != nil {
		return "", models.Session{}, err
	}

	return accessToken, session, nil
}
