package geminiclient

import "errors"

var (
	ErrGenerateMessage          = errors.New("failed to generate message")
	ErrUnexpectedResponseFormat = errors.New("unexpected response format")
)
