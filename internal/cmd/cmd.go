// Package cmd defines the command types and options shared between
// the processor layer and the Redis adapter layer.
package cmd

import "time"

// RedisCmd represents a supported Redis command name.
type RedisCmd string

const (
	RedisSet    RedisCmd = "SET"
	RedisGet    RedisCmd = "GET"
	RedisDel    RedisCmd = "DEL"
	RedisExists RedisCmd = "EXISTS"
	RedisExpire RedisCmd = "EXPIRE"
)

// RedisCmdData carries a parsed Redis command ready to be executed by an adapter.
// Depending on the command, different fields are populated:
//   - SET uses Key, Value, and SetOpts.
//   - GET uses Key.
//   - DEL uses KeysForDeletion (one or more keys).
//   - EXISTS uses Key.
//   - EXPIRE uses Key (options TBD).
type RedisCmdData struct {
	Command         RedisCmd
	Key             string
	KeysForDeletion []string
	Value           interface{}
	SetOpts         SetOptions
}

// SetOptions holds optional flags for the SET command.
// TTL-related fields (EX, PX, EXAT, PXAT, KeepTTL) are mutually exclusive.
// NX and XX are mutually exclusive.
type SetOptions struct {
	NX      bool
	XX      bool
	EX      time.Duration
	PX      time.Duration
	EXAT    time.Time
	PXAT    time.Time
	KeepTTL bool
}

// RedisCmdOut holds the result of executing a Redis command.
type RedisCmdOut struct {
	Command RedisCmd
	Data    []byte
}
