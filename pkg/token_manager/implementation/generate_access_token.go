package tokenmanagerimpl

import (
	"github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"github.com/golang-jwt/jwt/v5"
	"log/slog"
	"strconv"
	"time"
)

func (t *tokenManager) GenerateAccessToken(employeeID uint32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(t.accessTokenTTL),
		},
		Subject: strconv.Itoa(int(employeeID)),
	})

	signed, err := token.SignedString(t.key)
	if err != nil {
		t.logger.Error("failed to sign access token", slog.String("error", err.Error()))
		return "", token_manager.ErrAccessToken
	}

	return signed, nil
}
