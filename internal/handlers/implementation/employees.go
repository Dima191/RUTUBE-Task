package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/go-chi/render"
	"net/http"
)

func (h *handler) Employees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := h.service.Employees(r.Context())
		if err != nil {
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, employees)
	}
}
