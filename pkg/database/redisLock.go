package database

import (
	"errors"
	"time"

	"github.com/rehylas/wx/pkg/common"
)

type RedisLock struct {
	Name   string
	Expiry time.Duration
}

//TryLock ...
func (l *RedisLock) TryLock() error {
	if ok, _ := GetRedisClient().SetNX(l.Name, 1, l.Expiry).Result(); !ok {
		return errors.New(common.ERROR_REDIS_LOCK_EXIST)
	}
	return nil
}

//Unlock ...
func (l *RedisLock) Unlock() error {
	return GetRedisClient().Del(l.Name).Err()
}
