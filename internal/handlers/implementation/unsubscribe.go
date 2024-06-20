package handlerimpl

import (
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *handler) Unsubscribe() http.HandlerFunc {
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

		targetIDStr := chi.URLParam(r, handlers.EmployeeIDURLParam)
		if targetIDStr == "" {
			http.Error(w, handlers.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		targetID, err := strconv.Atoi(targetIDStr)
		if err != nil {
			http.Error(w, handlers.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		if err = h.service.Unsubscribe(r.Context(), uint32(employeeID), uint32(targetID)); err != nil {
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
