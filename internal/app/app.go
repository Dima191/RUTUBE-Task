package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Dima191/RUTUBE-Task/internal/config"
	h "github.com/Dima191/RUTUBE-Task/internal/handlers"
	handlerimpl "github.com/Dima191/RUTUBE-Task/internal/handlers/implementation"
	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
)

type App struct {
	cfg      *config.Config
	serv     http.Server
	router   *chi.Mux
	handlers h.Handler

	sp *serviceProvider

	configPath string

	logger *slog.Logger
}

func (a *App) initConfig(_ context.Context) error {
	cfg, err := config.New(a.configPath)
	if err != nil {
		return err
	}
	a.cfg = cfg

	return nil
}

func (a *App) initRouter(_ context.Context) error {
	a.router = chi.NewRouter()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider(a.cfg.JWTSignedKey, a.cfg.AccessTokenTTL, a.cfg.RefreshTokenTTL, a.cfg.GeminiAPIKey, a.cfg.DBConnectionString, a.cfg.SmtpHost, a.cfg.SmtpPort, a.cfg.SenderEmail, a.cfg.SMTPPassword, a.logger)
	return nil
}

func (a *App) initHandlers(ctx context.Context) error {
	s, err := a.sp.EmployeeService(ctx)
	if err != nil {
		return err
	}

	a.handlers = handlerimpl.New(s, a.sp.tokenManager, a.logger)
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initRouter,
		a.initHandlers,
	}

	for _, f := range deps {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) runHTTPServer() error {
	if err := a.handlers.Register(a.router); err != nil {
		return err
	}

	a.serv = http.Server{
		Addr:         fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port),
		Handler:      a.router,
		ReadTimeout:  a.cfg.ReadTimeout,
		WriteTimeout: a.cfg.WriteTimeout,
		IdleTimeout:  a.cfg.IdleTimeout,
	}

	a.logger.Info("starting http server on port", a.cfg.Host, a.cfg.Port)

	return a.serv.ListenAndServe()
}

func (a *App) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		worker, err := a.sp.NotificationWorker(ctx)
		if err != nil {
			a.logger.Error("failed to initialize notification worker", slog.String("error", err.Error()))
			return err
		}

		if err = worker.Run(ctx); err != nil {
			a.logger.Error("failed to start notification worker", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := a.runHTTPServer(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			a.logger.Error("failed to start http server", slog.String("error", err.Error()))
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	if err := a.serv.Close(); err != nil {
		return err
	}

	if a.sp.repository != nil {
		a.sp.repository.CloseConnection()
	}

	return nil
}

func New(ctx context.Context, configPath string, logger *slog.Logger) (*App, error) {
	a := &App{
		configPath: configPath,
		logger:     logger,
	}

	logger.Info("initializing app")
	if err := a.initDeps(ctx); err != nil {
		logger.Error("failed to initialize dependencies")
		return nil, err
	}

	return a, nil
}
