package go_redis

import (
	"context"

	"redis-viewer/internal/cmd"
)

// Exists checks whether a key exists. Not yet implemented.
func (a *RedisAdapter) Exists(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	//val, err := a.client.Exists(ctx, key).Result()
	//if err != nil {
	//	return false, fmt.Errorf("%w: %w", ErrExists, err)
	//}

	// return val == 1, nil

	return cmd.RedisCmdOut{}, nil
}
