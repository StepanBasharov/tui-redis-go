package go_redis

import "errors"

// Sentinel errors returned by RedisAdapter methods.
var (
	ErrFailedConnect = errors.New("failed to connect to redis")
	ErrCloseConn     = errors.New("failed to close connection")
	ErrGetValue      = errors.New("failed to get value")
	ErrNotFound      = errors.New("key not found")
	ErrSetValue      = errors.New("failed to set value")
	ErrDelValue      = errors.New("failed to delete value")
	ErrExpire        = errors.New("failed to expire")
	ErrExists        = errors.New("failed to check existence")
)
