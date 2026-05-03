// Package go_redis implements the redis.AdapterRedisProtocol interface
// using the redis/go-redis/v9 client library.
package go_redis

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisAdapter wraps a go-redis Client and implements redis.AdapterRedisProtocol.
type RedisAdapter struct {
	client *redis.Client
	host   string
	port   int
	db     int
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

	return &RedisAdapter{
		client: client,
		host:   host,
		port:   port,
		db:     db,
	}, nil
}

func (a *RedisAdapter) GetRedisAdd() string {
	return fmt.Sprintf("%s:%d/%d", a.host, a.port, a.db)
}

// Close shuts down the underlying Redis connection pool.
func (a *RedisAdapter) Close() error {
	if err := a.client.Close(); err != nil {
		return fmt.Errorf("%w: %w", ErrCloseConn, err)
	}

	return nil
}
