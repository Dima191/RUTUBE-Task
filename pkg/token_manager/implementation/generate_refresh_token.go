package tokenmanagerimpl

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/Dima191/RUTUBE-Task/internal/models"
	"github.com/Dima191/RUTUBE-Task/pkg/token_manager"
	"log/slog"
	"time"
	"unsafe"
)

func (t *tokenManager) GenerateRefreshToken() (models.Session, error) {
	refreshToken := models.Session{}
	var token [32]byte
	_, err := rand.Read(unsafe.Slice(&token[0], len(token)))
	if err != nil {
		t.logger.Error("failed to generate refresh token", slog.String("error", err.Error()))
		return models.Session{}, token_manager.ErrRefreshToken
	}

	refreshToken.Token = hex.EncodeToString(unsafe.Slice(&token[0], len(token)))
	refreshToken.ExpiresAt = time.Now().Add(t.refreshTokenTTL)

	return refreshToken, nil
}
