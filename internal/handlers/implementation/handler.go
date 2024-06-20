package handlerimpl

import (
	"errors"
	"github.com/Dima191/RUTUBE-Task/internal/handlers"
	internalmiddlewares "github.com/Dima191/RUTUBE-Task/internal/middlewares"
	employeesrv "github.com/Dima191/RUTUBE-Task/internal/service/employee"
	tm "github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log/slog"
	"net/http"
)

type handler struct {
	service employeesrv.Service
	token   tm.TokenManager

	logger *slog.Logger
}

func (h *handler) Register(r http.Handler) error {
	router, ok := r.(*chi.Mux)
	if !ok {
		h.logger.Error("failed to convert http.Handler to chi.Mux")
		return errors.New("failed to convert http.Handler to chi.Mux")
	}

	middlewares := []func(http.Handler) http.Handler{
		middleware.RequestID,
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.URLFormat,
	}

	router.Use(middlewares...)

	router.Post(handlers.SignUpURL, h.SignUp())
	router.Post(handlers.SignInURL, h.SignIn())

	router.Get(handlers.EmployeesURL, h.Employees())
	router.Get(handlers.EmployeeURL, h.Employee())

	router.Post(handlers.RefreshURL, h.Refresh())

	router.Route("/", func(router chi.Router) {
		router.Use(internalmiddlewares.Authorization(h.token))

		router.Get(handlers.SubscriptionsURL, h.Subscriptions())

		router.Route(handlers.EmployeeURL, func(router chi.Router) {
			router.Post(handlers.SubscribeURL, h.Subscribe())
			router.Post(handlers.UnsubscribeURL, h.Unsubscribe())
		})

		router.Post(handlers.LogOutURL, h.LogOut())
	})

	return nil
}

func New(service employeesrv.Service, token tm.TokenManager, logger *slog.Logger) handlers.Handler {
	h := &handler{
		service: service,
		token:   token,
		logger:  logger,
	}

	return h
}
