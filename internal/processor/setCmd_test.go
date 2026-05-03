package processor

import (
	"errors"
	"testing"
	"time"

	"github.com/StepanBasharov/tui-redis-go/internal/cmd"
)

func TestBuildSetCmd(t *testing.T) {
	tests := []struct {
		name      string
		tokens    []string
		wantKey   string
		wantValue any
		wantOpts  cmd.SetOptions
		wantErr   error
	}{
		{
			name:      "key and value only",
			tokens:    []string{"mykey", "myvalue"},
			wantKey:   "mykey",
			wantValue: "myvalue",
			wantOpts:  cmd.SetOptions{},
		},
		{
			name:    "missing value",
			tokens:  []string{"mykey"},
			wantErr: ErrNotEnoughArgs,
		},
		{
			name:    "empty tokens",
			tokens:  []string{},
			wantErr: ErrNotEnoughArgs,
		},
		{
			name:      "with NX",
			tokens:    []string{"mykey", "myvalue", "NX"},
			wantKey:   "mykey",
			wantValue: "myvalue",
			wantOpts:  cmd.SetOptions{NX: true},
		},
		{
			name:      "with XX",
			tokens:    []string{"mykey", "myvalue", "XX"},
			wantKey:   "mykey",
			wantValue: "myvalue",
			wantOpts:  cmd.SetOptions{XX: true},
		},
		{
			name:      "with EX",
			tokens:    []string{"mykey", "myvalue", "EX", "60"},
			wantKey:   "mykey",
			wantValue: "myvalue",
			wantOpts:  cmd.SetOptions{EX: 60 * time.Second},
		},
		{
			name:      "with NX and EX",
			tokens:    []string{"mykey", "myvalue", "NX", "EX", "30"},
			wantKey:   "mykey",
			wantValue: "myvalue",
			wantOpts:  cmd.SetOptions{NX: true, EX: 30 * time.Second},
		},
		{
			name:    "EX without value",
			tokens:  []string{"mykey", "myvalue", "EX"},
			wantErr: ErrInvalidSetOptions,
		},
		{
			name:    "EX with non-numeric value",
			tokens:  []string{"mykey", "myvalue", "EX", "abc"},
			wantErr: ErrInvalidSetOptions,
		},
		{
			name:    "unknown option",
			tokens:  []string{"mykey", "myvalue", "INVALID"},
			wantErr: ErrInvalidSetOptions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rCmd := cmd.RedisCmdData{Command: cmd.RedisSet}

			err := buildSetCmd(&rCmd, tt.tokens)

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
			if rCmd.Value != tt.wantValue {
				t.Errorf("Value = %v, want %v", rCmd.Value, tt.wantValue)
			}
			if rCmd.SetOpts != tt.wantOpts {
				t.Errorf("SetOpts = %+v, want %+v", rCmd.SetOpts, tt.wantOpts)
			}
		})
	}
}

func TestParseSetOptions(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []string
		want    cmd.SetOptions
		wantErr error
	}{
		{
			name:   "empty tokens",
			tokens: []string{},
			want:   cmd.SetOptions{},
		},
		{
			name:   "NX only",
			tokens: []string{"NX"},
			want:   cmd.SetOptions{NX: true},
		},
		{
			name:   "XX only",
			tokens: []string{"XX"},
			want:   cmd.SetOptions{XX: true},
		},
		{
			name:   "EX with seconds",
			tokens: []string{"EX", "120"},
			want:   cmd.SetOptions{EX: 120 * time.Second},
		},
		{
			name:    "EX missing value",
			tokens:  []string{"EX"},
			wantErr: ErrInvalidSetOptions,
		},
		{
			name:    "EX invalid number",
			tokens:  []string{"EX", "notanumber"},
			wantErr: ErrInvalidSetOptions,
		},
		{
			name:    "unknown flag",
			tokens:  []string{"FOOBAR"},
			wantErr: ErrInvalidSetOptions,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSetOptions(tt.tokens)

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

			if got != tt.want {
				t.Errorf("parseSetOptions() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
