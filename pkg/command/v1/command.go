package v1

import "context"

type Command interface {
	execute(ctx context.Context)
}
