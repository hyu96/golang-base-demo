package credis

import (
	redis "github.com/go-redis/redis/v8"
)

var (
	ErrNil    = redis.Nil
	ErrClosed = redis.ErrClosed
)
