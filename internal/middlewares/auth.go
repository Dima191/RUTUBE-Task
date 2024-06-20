package middlewares

import (
	"context"
	"errors"
	tm "github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

const (
	RefreshTokenCookiesKey = "refresh_token"
	AuthorizationHeader    = "Authorization"
	EmployeeIDKeyContext   = "employee_id"
)

func Authorization(token tm.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get(AuthorizationHeader)
			headerParts := strings.Split(header, " ")
			if len(headerParts) != 2 && headerParts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			claims, err := token.Parse(headerParts[1])
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					http.Error(w, "access token expired. refresh it", http.StatusUnauthorized)
					return
				}

				http.Error(w, "fail to get access token", http.StatusUnauthorized)
				return
			}

			employeeID, err := claims.GetSubject()
			if err != nil {
				http.Error(w, "fail to get access token", http.StatusUnauthorized)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), EmployeeIDKeyContext, employeeID))
			next.ServeHTTP(w, r)
		})
	}
}
