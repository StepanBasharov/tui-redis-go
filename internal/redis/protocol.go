// Package redis defines the port interface for Redis operations.
// Concrete adapters (e.g. go-redis) implement this interface.
package redis

import (
	"context"

	"redis-viewer/internal/cmd"
)

// AdapterRedisProtocol is the port that any Redis client adapter must satisfy.
// All methods accept a fully-populated RedisCmdData with the relevant fields
// set for the given command.
type AdapterRedisProtocol interface {
	GetValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error)
	SetValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error)
	DelValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error)
	Expire(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error)
	Exists(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error)
	GetRedisAdd() string
	Close() error
}
