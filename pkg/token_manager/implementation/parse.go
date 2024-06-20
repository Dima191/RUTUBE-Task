package tokenmanagerimpl

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func (t *tokenManager) Parse(accessToken string) (jwt.Claims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return t.key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
