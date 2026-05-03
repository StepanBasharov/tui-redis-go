package processor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"redis-viewer/internal/cmd"
)

func (p *Processor) processSet(ctx context.Context, rCmd cmd.RedisCmdData, tokens []string) (cmd.RedisCmdOut, error) {
	if err := buildSetCmd(&rCmd, tokens); err != nil {
		return cmd.RedisCmdOut{}, err
	}

	out, err := p.client.SetValue(ctx, rCmd)
	if err != nil {
		return cmd.RedisCmdOut{}, err
	}

	return out, nil
}

// buildSetCmd populates rCmd with key, value, and parsed SET options from tokens.
// Expects at least two tokens: key and value. Remaining tokens are parsed as options.
func buildSetCmd(rCmd *cmd.RedisCmdData, tokens []string) error {
	if len(tokens) < 2 {
		return fmt.Errorf("%w: SET requires key and value", ErrNotEnoughArgs)
	}

	rCmd.Key = tokens[0]
	rCmd.Value = tokens[1]

	opts, err := parseSetOptions(tokens[2:])
	if err != nil {
		return err
	}

	rCmd.SetOpts = opts

	return nil
}

// parseSetOptions iterates over tokens after key/value and populates SetOptions.
// Currently supports NX, XX, and EX <seconds>.
// Returns ErrInvalidSetOptions on unknown flags or malformed values.
func parseSetOptions(tokens []string) (cmd.SetOptions, error) {
	var opts cmd.SetOptions

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "NX":
			opts.NX = true
		case "XX":
			opts.XX = true
		case "EX":
			i++
			if i >= len(tokens) {
				return opts, fmt.Errorf("%w: EX requires a value", ErrInvalidSetOptions)
			}
			dur, err := strconv.Atoi(tokens[i])
			if err != nil {
				return opts, fmt.Errorf("%w: invalid EX value: %w", ErrInvalidSetOptions, err)
			}
			opts.EX = time.Duration(dur) * time.Second
		default:
			return opts, fmt.Errorf("%w: %s", ErrInvalidSetOptions, tokens[i])
		}
	}

	return opts, nil
}
