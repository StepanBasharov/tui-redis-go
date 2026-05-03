package ui

import (
	"context"
	"strconv"

	"redis-viewer/internal/redis"
	go_redis "redis-viewer/internal/redis/go-redis"
)

type History struct {
	cmdHist string
	resHist string
}

type Connection struct {
	Client  redis.AdapterRedisProtocol
	History []History
}

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
