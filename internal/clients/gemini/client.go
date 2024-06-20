package geminiclient

import "context"

type Client interface {
	GenerateMessage(ctx context.Context, query string) (string, error)
}
