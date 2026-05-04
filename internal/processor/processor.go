// Package processor parses raw user input into typed Redis commands
// and dispatches them to a redis.AdapterRedisProtocol implementation.
package processor

import (
	"context"
	"fmt"
	"strings"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"

	"github.com/StepanBasharov/tui-redis-go/internal/redis"
)

// Processor holds a reference to the Redis adapter and exposes ProcessCmd
// as the single entry point for command execution.
type Processor struct {
	client redis.AdapterRedisProtocol
}

// NewProcessor creates a Processor backed by the given adapter.
func NewProcessor(client redis.AdapterRedisProtocol) *Processor {
	return &Processor{
		client: client,
	}
}

// ProcessCmd tokenizes a raw command string, identifies the Redis command,
// and delegates to the appropriate handler (processSet, processGet, etc.).
func (p *Processor) ProcessCmd(ctx context.Context, command string) cmd.RedisCmdOut {
	cmdTokens := strings.Split(command, " ")

	if len(cmdTokens) == 0 {
		return cmd.RedisCmdOut{Data: []byte("Empty Command")}
	}

	cmdData, err := p.pickCommand(cmdTokens[0])
	if err != nil {
		return cmd.RedisCmdOut{Data: []byte(err.Error())}
	}

	cmdTokens = cmdTokens[1:]

	switch cmdData.Command {
	case cmd.RedisSet:
		out, err := p.processSet(ctx, cmdData, cmdTokens)
		if err != nil {
			return cmd.RedisCmdOut{Data: []byte(err.Error())}
		}

		return out

	case cmd.RedisGet:
		out, err := p.processGet(ctx, cmdData, cmdTokens)
		if err != nil {
			return cmd.RedisCmdOut{Data: []byte(err.Error())}
		}

		return out

	case cmd.RedisDel:
		out, err := p.processDel(ctx, cmdData, cmdTokens)
		if err != nil {
			return cmd.RedisCmdOut{Data: []byte(err.Error())}
		}

		return out

	case cmd.RedisExists:
		out, err := p.processExists(ctx, cmdData, cmdTokens)
		if err != nil {
			return cmd.RedisCmdOut{Data: []byte(err.Error())}
		}

		return out

	case cmd.RedisExpire:
		if err := p.processExpire(ctx, cmdData, cmdTokens); err != nil {
			return cmd.RedisCmdOut{Data: []byte(err.Error())}
		}
	}

	return cmd.RedisCmdOut{}
}

// pickCommand maps the first token of user input to a RedisCmdData
// with the Command field set. Returns an error for unknown commands.
func (p *Processor) pickCommand(firstToken string) (cmd.RedisCmdData, error) {
	switch strings.ToUpper(firstToken) {
	case "SET":
		return cmd.RedisCmdData{Command: cmd.RedisSet}, nil
	case "GET":
		return cmd.RedisCmdData{Command: cmd.RedisGet}, nil
	case "DEL":
		return cmd.RedisCmdData{Command: cmd.RedisDel}, nil
	case "EXISTS":
		return cmd.RedisCmdData{Command: cmd.RedisExists}, nil
	case "EXPIRE":
		return cmd.RedisCmdData{Command: cmd.RedisExpire}, nil
	default:
		return cmd.RedisCmdData{}, fmt.Errorf("unknown command: %s", firstToken)
	}
}

// Close shuts down the underlying Redis adapter connection.
func (p *Processor) Close() error {
	return p.client.Close()
}
