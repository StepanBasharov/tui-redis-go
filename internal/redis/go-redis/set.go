package go_redis

import (
	"context"
	"fmt"

	"redis-viewer/internal/cmd"

	"github.com/redis/go-redis/v9"
)

// SetValue executes a SET command with all options from data.SetOpts
// (NX, XX, EX, PX, EXAT, PXAT, KeepTTL) mapped to redis.SetArgs.
func (a *RedisAdapter) SetValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	opts := data.SetOpts

	setArgs := redis.SetArgs{}

	switch {
	case opts.EX > 0:
		setArgs.TTL = opts.EX
	case opts.PX > 0:
		setArgs.TTL = opts.PX
	case !opts.EXAT.IsZero():
		setArgs.ExpireAt = opts.EXAT
	case !opts.PXAT.IsZero():
		setArgs.ExpireAt = opts.PXAT
	case opts.KeepTTL:
		setArgs.KeepTTL = true
	}

	if opts.NX {
		setArgs.Mode = "NX"
	} else if opts.XX {
		setArgs.Mode = "XX"
	}

	res, err := a.client.SetArgs(ctx, data.Key, data.Value, setArgs).Result()
	if err != nil {
		fmt.Println(err.Error())
		return cmd.RedisCmdOut{}, fmt.Errorf("%w: %w", ErrSetValue, err)
	}

	return cmd.RedisCmdOut{Data: []byte(res)}, nil
}
