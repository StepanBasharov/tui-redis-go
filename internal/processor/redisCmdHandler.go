package processor

import (
	"context"

	"redis-viewer/internal/cmd"
)

func (p *Processor) processExists(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) error {
	return nil
}

func (p *Processor) processExpire(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) error {
	return nil
}
