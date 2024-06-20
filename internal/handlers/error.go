package handlers

import "errors"

var (
	ErrInternal                = errors.New("internal error")
	ErrInvalidEmployeeID       = errors.New("invalid employee id")
	ErrEmployeeNotFound        = errors.New("employee not found")
	ErrNoRefreshTokenInCookies = errors.New("no refresh token found in cookies")
	ErrRefreshTokenExpired     = errors.New("refresh token is expired")
	ErrInvalidRequestBody      = errors.New("invalid request body. unable to decode")
)
