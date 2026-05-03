package processor

import (
	"context"
	"fmt"

	"redis-viewer/internal/cmd"
)

func (p *Processor) processGet(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) (cmd.RedisCmdOut, error) {
	if err := buildGetCmd(&rCmd, tokens); err != nil {
		return cmd.RedisCmdOut{}, err
	}

	out, err := p.client.GetValue(ctx, rCmd)
	if err != nil {
		return cmd.RedisCmdOut{}, err
	}

	return out, nil
}

// buildGetCmd populates rCmd with the key from tokens.
// Expects at least one token. Extra tokens are ignored.
func buildGetCmd(rCmd *cmd.RedisCmdData, tokens []string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%w: GET requires key", ErrNotEnoughArgs)
	}

	rCmd.Key = tokens[0]

	return nil
}
