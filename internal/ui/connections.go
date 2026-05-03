package ui

import (
	"context"
	"strconv"

	"redis-viewer/internal/redis"
	go_redis "redis-viewer/internal/redis/go-redis"
)

// History records a single command and its response.
type History struct {
	cmdHist string
	resHist string
}

// Connection pairs a Redis adapter with its command history.
type Connection struct {
	Client  redis.AdapterRedisProtocol
	History []History
}

// NewConnection dials a Redis instance and returns a Connection.
func NewConnection(host, port, password, database string) (*Connection, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	dbInt, err := strconv.Atoi(database)
	if err != nil {
		return nil, err
	}

	client, err := go_redis.NewRedisAdapter(context.Background(), host, portInt, password, dbInt)
	if err != nil {
		return nil, err
	}

	history := make([]History, 0)

	return &Connection{client, history}, nil
}
