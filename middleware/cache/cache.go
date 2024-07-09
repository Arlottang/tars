package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type cache struct {
	config *Conf

	runtime *sync.Map
	cli     *redis.Client
}

func NewCache(cfg *Conf) ISyncCache {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       0,
	})

	return &cache{
		config:  cfg,
		runtime: new(sync.Map),
		cli:     cli,
	}
}

func (c *cache) Lock() {
	c.LockWithKey(defaultGlobalLockKey)

}

func (c *cache) Unlock() {
	c.UnlockWithKey(defaultGlobalLockKey)
}

func (c *cache) LockWithKey(key any) {
	ctx := context.Background()
	lockKey := defaultLockKey + c.config.Topic + fmt.Sprintf("%v", key)

	for {
		ok, err := c.cli.SetNX(ctx, lockKey, "locked", c.config.LockTime).Result()
		if err != nil {
			//w.log.Errorf("cache lock panic, err: %v", err)
			return
		}

		if ok {
			break
		}

		time.Sleep(time.Second)
	}

}

func (c *cache) UnlockWithKey(key any) {
	ctx := context.Background()
	lockKey := defaultLockKey + c.config.Topic + fmt.Sprintf("%v", key)

	for {
		_, err := c.cli.Del(ctx, lockKey).Result()
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	return

}

func (c *cache) Set(key, value any) {
	c.runtime.Store(key, value)
}

func (c *cache) Get(key any) (any, error) {
	value, ok := c.runtime.Load(key)
	if !ok {
		return nil, errors.New(fmt.Sprintf("unable to load flowinstance with %v", key))
	}

	return value, nil

}

func (c *cache) Del(key any) {
	c.runtime.Delete(key)
}

func (c *cache) Range(f func(key any, value any) bool) {
	c.runtime.Range(f)
}

func (c *cache) Persist(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c *cache) Recover(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
