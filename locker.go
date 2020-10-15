package cron

import "time"

type Locker interface {
	Obtain(key string, ttl time.Duration) (Lock, error)
}

type Lock interface {
	TTL() (time.Duration, error)
	Release() error
}
