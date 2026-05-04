package go_redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

// Exists checks whether a key exists and returns the count of existing keys.
func (a *RedisAdapter) Exists(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	res, err := a.client.Exists(ctx, data.Keys...).Result()
	if err != nil {
		return cmd.RedisCmdOut{}, fmt.Errorf("%w: %w", ErrExists, err)
	}

	out := strconv.Itoa(int(res))

	return cmd.RedisCmdOut{Data: []byte(out)}, nil
}
