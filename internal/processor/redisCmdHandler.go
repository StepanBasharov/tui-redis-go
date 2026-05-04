package processor

import (
	"context"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

// processExpire is a stub handler for the EXPIRE command. Not yet implemented.
func (p *Processor) processExpire(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) error {
	return nil
}
