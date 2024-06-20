package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/render"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func (h *handler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		credentials := models.SignIn{}
		if err := render.DecodeJSON(r.Body, &credentials); err != nil {
			http.Error(w, handlers.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
			return
		}

		accessToken, refreshToken, err := h.service.SignIn(r.Context(), credentials)
		if err != nil {
			var validationErr *multierror.Error
			switch {
			case errors.Is(err, srv.ErrInvalidLoginOrPassword):
				http.Error(w, srv.ErrInvalidLoginOrPassword.Error(), http.StatusBadRequest)
				return
			case errors.As(err, &validationErr):
				http.Error(w, err.Error(), http.StatusBadRequest)
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

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{accessToken, refreshToken})
	}
}
