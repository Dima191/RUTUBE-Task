package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/go-chi/render"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func (h *handler) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employee := models.SignUp{}
		if err := render.DecodeJSON(r.Body, &employee); err != nil {
			if errors.Is(err, models.ErrInvalidDateFormat) {
				http.Error(w, models.ErrInvalidDateFormat.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, handlers.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
			return
		}

		accessToken, refreshToken, err := h.service.SignUp(r.Context(), employee)
		if err != nil {
			var validationErr *multierror.Error
			switch {
			case errors.As(err, &validationErr):
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			case errors.Is(err, rep.ErrEmployeeAlreadyExists):
				http.Error(w, rep.ErrEmployeeAlreadyExists.Error(), http.StatusBadRequest)
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

		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		}{accessToken, refreshToken})
	}
}
