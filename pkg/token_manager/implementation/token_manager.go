package tokenmanagerimpl

import (
	"github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"log/slog"
	"time"
)

type tokenManager struct {
	key             []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	logger *slog.Logger
}

func New(key []byte, accessTokenTTL, refreshTokenTTL time.Duration, logger *slog.Logger) token_manager.TokenManager {
	return &tokenManager{
		key:             key,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		logger:          logger,
	}
}
