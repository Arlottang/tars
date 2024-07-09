package cache

import (
	"context"
)

type ISyncCache interface {
	ISyncLock
	ISyncMap

	Persist(ctx context.Context) error
	Recover(ctx context.Context) error
}

type ISyncLock interface {
	Lock()
	Unlock()

	LockWithKey(key any)
	UnlockWithKey(key any)
}

type ISyncMap interface {
	Set(key, value any)
	Get(key any) (any, error)
	Del(key any)

	Range(f func(key, value any) bool)
}
