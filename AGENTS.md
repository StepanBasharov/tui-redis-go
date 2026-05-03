# AGENTS.md

## Project

Interactive Redis CLI viewer. Go 1.25, module `redis-viewer`.

## Architecture

- `cmd/redis-viewer/rv.go` -- entrypoint (bubbletea TUI). `main.go` at root is a GoLand template placeholder, not the app.
- `internal/ui/` -- bubbletea Model/Update/View split across `model.go`, `update.go`, `view.go`. Never use `fmt.Print*` inside `View()` -- it corrupts bubbletea's terminal rendering.
- `internal/ui/ui-components/` -- stateless rendering functions (e.g. `banners.go`, `addNewConnection.go`). Centered via `lipgloss.Place` using terminal width/height from `tea.WindowSizeMsg`.
- `internal/ui/forms/` -- form structs with their own `Update`/`CycleFocus`/`ResetFocus` methods (e.g. `ConnectionForm`). Forms hold named `textinput.Model` fields (not arrays with index constants). The parent `Update()` must forward `tea.Msg` into the active form's `Update()` for inputs to work.
- `internal/redis/protocol.go` -- `AdapterRedisProtocol` interface (the port)
- `internal/redis/go-redis/` -- adapter implementation using `redis/go-redis/v9`
- `internal/cmd/` -- command types (`RedisCmdData`, `SetOptions`)
- `internal/processor/` -- parses user input into `RedisCmdData` and calls the adapter. One file per command: `setCmd.go`, `getCmd.go`, `delCmd.go`, etc. Tests follow the same split: `setCmd_test.go`, `getCmd_test.go`, `delCmd_test.go`.

## Commands

```bash
# run the CLI
go run ./cmd/redis-viewer

# tests (stdlib only, no testify)
go test ./...

# single package
go test -v ./internal/processor/
```

## Conventions

- Tests use stdlib `testing` only -- no testify
- Table-driven tests with `t.Run` subtests
- Adapter methods accept `cmd.RedisCmdData` (not individual args)
- Each command handler has a `build*Cmd(*cmd.RedisCmdData, []string) error` function that fills the command struct, separate from `process*` which calls the client
- SET-specific options live in `SetOptions`, not a generic `Options` struct. Field on `RedisCmdData` is `SetOpts`.
- DEL stores all keys in `RedisCmdData.KeysForDeletion` (not `Key`/`Value`)
- Sentinel errors per package (e.g. `ErrNotEnoughArgs`, `ErrInvalidSetOptions` in processor)
- Constructor `NewRedisAdapter` takes `context.Context` as first param
