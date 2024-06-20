package notification

import "context"

type Service interface {
	Notice(ctx context.Context, subscriberFullName, subscriberEmail string, celebrantFullName, celebrantEmail string) error
}
