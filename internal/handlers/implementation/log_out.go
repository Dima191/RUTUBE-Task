package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	"net/http"
)

func (h *handler) LogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh, err := r.Cookie(middlewares.RefreshTokenCookiesKey)
		if err != nil {
			http.Error(w, handlers.ErrNoRefreshTokenInCookies.Error(), http.StatusUnauthorized)
			return
		}

		if err = h.service.LogOut(r.Context(), refresh.Value); err != nil {
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	}
}
