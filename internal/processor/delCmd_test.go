package processor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

func TestBuildDelCmd(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		wantKeys []string
		wantErr  error
	}{
		{
			name:     "single key",
			tokens:   []string{"mykey"},
			wantKeys: []string{"mykey"},
		},
		{
			name:    "empty tokens",
			tokens:  []string{},
			wantErr: ErrNotEnoughArgs,
		},
		{
			name:     "multiple keys",
			tokens:   []string{"key1", "key2", "key3"},
			wantKeys: []string{"key1", "key2", "key3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rCmd := cmd.RedisCmdData{Command: cmd.RedisDel}

			err := buildDelCmd(&rCmd, tt.tokens)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if fmt.Sprintf("%v", rCmd.Keys) != fmt.Sprintf("%v", tt.wantKeys) {
				t.Errorf("Keys = %v, want %v", rCmd.Keys, tt.wantKeys)
			}
		})
	}
}
