package processor

import (
	"context"
	"fmt"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

// processExists builds and executes an EXISTS command from the given tokens.
func (p *Processor) processExists(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) (cmd.RedisCmdOut, error) {
	if err := buildExistsCmd(&rCmd, tokens); err != nil {
		return cmd.RedisCmdOut{}, err
	}

	out, err := p.client.Exists(ctx, rCmd)
	if err != nil {
		return cmd.RedisCmdOut{}, err
	}

	return out, nil
}

// buildExistsCmd populates rCmd.Key with the key to check.
// Expects exactly one token (key).
func buildExistsCmd(rCmd *cmd.RedisCmdData, tokens []string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%w: EXISTS requires at least one key", ErrNotEnoughArgs)
	}

	rCmd.Keys = make([]string, 0, len(tokens))
	rCmd.Keys = append(rCmd.Keys, tokens...)

	return nil
}
