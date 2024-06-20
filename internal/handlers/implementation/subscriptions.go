package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *handler) Subscriptions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr, ok := r.Context().Value(middlewares.EmployeeIDKeyContext).(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		employeeID, err := strconv.Atoi(employeeIDStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		employees, err := h.service.Subscriptions(r.Context(), uint32(employeeID))
		if err != nil {
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, employees)
	}
}
