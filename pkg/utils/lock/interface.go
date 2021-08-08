package lock

import (
	"context"
)

type Interface interface {
	Lock(ctx context.Context, key string) error
	UnLock(ctx context.Context, key string) error
}
