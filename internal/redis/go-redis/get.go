package go_redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"

	"github.com/redis/go-redis/v9"
)

// GetValue executes a GET command. Returns ErrNotFound if the key does not exist.
func (a *RedisAdapter) GetValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	val, err := a.client.Get(ctx, data.Key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return cmd.RedisCmdOut{}, ErrNotFound
		}

		return cmd.RedisCmdOut{}, fmt.Errorf("%w: %w", ErrGetValue, err)
	}

	out := []byte(val)

	return cmd.RedisCmdOut{Data: out}, nil
}
