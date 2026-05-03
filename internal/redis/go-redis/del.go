package go_redis

import (
	"context"
	"fmt"
	"strconv"

	"redis-viewer/internal/cmd"
)

// DelValue executes a DEL command for one or more keys from data.KeysForDeletion.
func (a *RedisAdapter) DelValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	res, err := a.client.Del(ctx, data.KeysForDeletion...).Result()
	if err != nil {
		return cmd.RedisCmdOut{}, fmt.Errorf("%w: %w", ErrDelValue, err)
	}

	out := strconv.Itoa(int(res))

	return cmd.RedisCmdOut{Data: []byte(out)}, nil
}
