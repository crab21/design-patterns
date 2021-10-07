package v1

import "context"

type Device interface {
	on(ctx context.Context)
	off(ctx context.Context)
}
