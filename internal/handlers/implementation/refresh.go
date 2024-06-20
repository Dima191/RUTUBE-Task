package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/render"
	"net/http"
)

func (h *handler) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh, err := r.Cookie(middlewares.RefreshTokenCookiesKey)
		if err != nil {
			http.Error(w, handlers.ErrNoRefreshTokenInCookies.Error(), http.StatusUnauthorized)
			return
		}

		accessToken, refreshToken, err := h.service.UpdateTokens(r.Context(), refresh.Value)
		if err != nil {
			switch {
			case errors.Is(err, rep.ErrNoSession):
				http.Error(w, rep.ErrNoSession.Error(), http.StatusUnauthorized)
				return
			case errors.Is(err, srv.ErrTokenExpired):
				http.Error(w, handlers.ErrRefreshTokenExpired.Error(), http.StatusUnauthorized)
				return
			default:
				http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.SetCookie(w, &http.Cookie{
			Name:     middlewares.RefreshTokenCookiesKey,
			Path:     "/",
			Value:    refreshToken,
			HttpOnly: true,
			Secure:   true,
		})

		render.JSON(w, r, struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{accessToken, refreshToken})
	}
}
