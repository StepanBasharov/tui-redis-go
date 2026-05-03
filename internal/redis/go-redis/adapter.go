// Package go_redis implements the redis.AdapterRedisProtocol interface
// using the redis/go-redis/v9 client library.
package go_redis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"redis-viewer/internal/cmd"

	"github.com/redis/go-redis/v9"
)

// RedisAdapter wraps a go-redis Client and implements redis.AdapterRedisProtocol.
type RedisAdapter struct {
	client *redis.Client
}

// NewRedisAdapter connects to a Redis instance and verifies the connection with PING.
// The client is closed if the health check fails to prevent resource leaks.
func NewRedisAdapter(ctx context.Context, host string, port int, password string, db int) (*RedisAdapter, error) {
	portToString := strconv.Itoa(port)
	addr := net.JoinHostPort(host, portToString)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		if errClose := client.Close(); errClose != nil {
			return nil, errors.Join(fmt.Errorf("%w: %w", ErrFailedConnect, err), ErrCloseConn)
		}

		return nil, fmt.Errorf("%w: %w", ErrFailedConnect, err)
	}

	return &RedisAdapter{client: client}, nil
}

// Close shuts down the underlying Redis connection pool.
func (a *RedisAdapter) Close() error {
	if err := a.client.Close(); err != nil {
		return fmt.Errorf("%w: %w", ErrCloseConn, err)
	}

	return nil
}

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

// DelValue executes a DEL command for one or more keys from data.KeysForDeletion.
func (a *RedisAdapter) DelValue(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	res, err := a.client.Del(ctx, data.KeysForDeletion...).Result()
	if err != nil {
		return cmd.RedisCmdOut{}, fmt.Errorf("%w: %w", ErrDelValue, err)
	}

	out := strconv.Itoa(int(res))

	return cmd.RedisCmdOut{Data: []byte(out)}, nil
}

// Expire sets a TTL on a key. Not yet implemented.
func (a *RedisAdapter) Expire(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	//if err := a.client.Expire(ctx, key, duration).Err(); err != nil {
	//	return fmt.Errorf("%w: %w", ErrExpire, err)
	//}

	return cmd.RedisCmdOut{}, nil
}

// Exists checks whether a key exists. Not yet implemented.
func (a *RedisAdapter) Exists(ctx context.Context, data cmd.RedisCmdData) (cmd.RedisCmdOut, error) {
	//val, err := a.client.Exists(ctx, key).Result()
	//if err != nil {
	//	return false, fmt.Errorf("%w: %w", ErrExists, err)
	//}

	// return val == 1, nil

	return cmd.RedisCmdOut{}, nil
}
