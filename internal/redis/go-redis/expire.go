package go_redis

import (
	"context"
	"redis-viewer/internal/cmd"
)

// Expire sets a TTL on a key. Not yet implemented.
func (a *RedisAdapter) Expire(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	//if err := a.client.Expire(ctx, key, duration).Err(); err != nil {
	//	return fmt.Errorf("%w: %w", ErrExpire, err)
	//}

	return cmd.RedisCmdOut{}, nil
}
