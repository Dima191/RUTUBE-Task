package token_manager

import "errors"

var (
	ErrAccessToken  = errors.New("access token error")
	ErrRefreshToken = errors.New("refresh token error")
)
