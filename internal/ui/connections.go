package ui

import (
	"context"
	"strconv"

	"github.com/StepanBasharov/tui-redis-go/internal/redis"
	go_redis "github.com/StepanBasharov/tui-redis-go/internal/redis/go-redis"
)

// ConnectionItem adapts a connection address for use with bubbles/list.
type ConnectionItem struct {
	Addr string
}

// Title returns the connection address shown in the list.
func (i ConnectionItem) Title() string { return i.Addr }

// Description returns an empty string (single-line list items).
func (i ConnectionItem) Description() string { return "" }

// FilterValue returns the string used by the list's built-in filter.
func (i ConnectionItem) FilterValue() string { return i.Addr }

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
