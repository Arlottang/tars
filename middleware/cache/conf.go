package cache

import "time"

type Conf struct {
	Topic    string
	LockTime time.Duration

	*RedisConf
}

type RedisConf struct {
	Address  string
	Password string
}
