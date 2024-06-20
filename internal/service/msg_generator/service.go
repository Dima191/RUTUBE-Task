package msggenerator

import "context"

type Service interface {
	Generate(ctx context.Context, subscriberFullName, celebrantFullName string) (string, error)
}
