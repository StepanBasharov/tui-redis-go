package processor

import (
	"context"
	"fmt"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

// processDel builds and executes a DEL command from the given tokens.
func (p *Processor) processDel(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) (cmd.RedisCmdOut, error) {
	if err := buildDelCmd(&rCmd, tokens); err != nil {
		return cmd.RedisCmdOut{}, err
	}

	out, err := p.client.DelValue(ctx, rCmd)
	if err != nil {
		return cmd.RedisCmdOut{}, err
	}

	return out, nil
}

// buildDelCmd populates rCmd.KeysForDeletion with all tokens.
// Expects at least one token (key). Supports multiple keys for batch deletion.
func buildDelCmd(rCmd *cmd.RedisCmdData, tokens []string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%w: DEL requires at least one key", ErrNotEnoughArgs)
	}

	rCmd.KeysForDeletion = make([]string, 0, len(tokens))

	rCmd.KeysForDeletion = append(rCmd.KeysForDeletion, tokens...)

	return nil
}
