package postgresrep

import (
	"context"
	rep "github.com/Dima191/RUTUBE-Task/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type repository struct {
	pool *pgxpool.Pool

	logger *slog.Logger
}

func New(ctx context.Context, connectionStr string, logger *slog.Logger) (rep.Repository, error) {
	pool, err := pgxpool.New(ctx, connectionStr)
	if err != nil {
		logger.Error("failed to connect to database", slog.String("error", err.Error()))
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		logger.Error("failed to ping database", slog.String("error", err.Error()))
		return nil, err
	}

	return &repository{
		pool:   pool,
		logger: logger,
	}, nil
}
