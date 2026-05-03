package processor

import (
	"errors"
	"testing"

	"redis-viewer/internal/cmd"
)

func TestBuildGetCmd(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []string
		wantKey string
		wantErr error
	}{
		{
			name:    "valid key",
			tokens:  []string{"mykey"},
			wantKey: "mykey",
		},
		{
			name:    "empty tokens",
			tokens:  []string{},
			wantErr: ErrNotEnoughArgs,
		},
		{
			name:    "extra tokens ignored",
			tokens:  []string{"mykey", "extra"},
			wantKey: "mykey",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rCmd := cmd.RedisCmdData{Command: cmd.RedisGet}

			err := buildGetCmd(&rCmd, tt.tokens)

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

			if rCmd.Key != tt.wantKey {
				t.Errorf("Key = %q, want %q", rCmd.Key, tt.wantKey)
			}
		})
	}
}
