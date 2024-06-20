package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	"github.com/Dima191/RUTUBE-Task/internal/middlewares"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	srv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func (h *handler) Subscribe() http.HandlerFunc {
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
			http.Error(w, srv.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		targetID, err := strconv.Atoi(targetIDStr)
		if err != nil {
			http.Error(w, srv.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		if err = h.service.Subscribe(r.Context(), uint32(employeeID), uint32(targetID)); err != nil {
			switch {
			case errors.Is(err, srv.ErrSelfSubscription):
				http.Error(w, srv.ErrSelfSubscription.Error(), http.StatusBadRequest)
				return
			case errors.Is(err, rep.ErrAlreadySubscribed):
				http.Error(w, rep.ErrAlreadySubscribed.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
