package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (h *handler) Employee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr := chi.URLParam(r, handlers.EmployeeIDURLParam)
		if employeeIDStr == "" {
			http.Error(w, handlers.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		employeeID, err := strconv.Atoi(employeeIDStr)
		if err != nil {
			http.Error(w, handlers.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}

		employee, err := h.service.EmployeeByID(r.Context(), uint32(employeeID))
		if err != nil {
			if errors.Is(err, rep.ErrEmployeeNotFound) {
				http.Error(w, handlers.ErrEmployeeNotFound.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, handlers.ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, employee)
	}
}
