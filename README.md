# redis-viewer

Interactive Redis CLI with syntax highlighting and command autocompletion.

## Requirements

- Go 1.25+
- A running Redis instance

## Installation

```bash
go install ./cmd/redis-viewer
```

Or run directly:

```bash
go run ./cmd/redis-viewer
```

## Supported commands

| Command  | Syntax                              | Description                       |
|----------|-------------------------------------|-----------------------------------|
| `SET`    | `SET key value [NX\|XX] [EX sec]`  | Set a key with optional flags     |
| `GET`    | `GET key`                           | Get the value of a key            |
| `DEL`    | `DEL key [key ...]`                 | Delete one or more keys           |
| `EXISTS` | `EXISTS key`                        | Check if a key exists             |
| `EXPIRE` | `EXPIRE key seconds`               | Set a TTL on a key                |

### SET options

- `NX` -- only set if the key does not already exist
- `XX` -- only set if the key already exists
- `EX <seconds>` -- set expiration in seconds
- `PX <milliseconds>` -- set expiration in milliseconds
- `EXAT <timestamp>` -- set expiration as Unix timestamp (seconds)
- `PXAT <timestamp>` -- set expiration as Unix timestamp (milliseconds)
- `KEEPTTL` -- retain the existing TTL when overwriting

## Project structure

```
cmd/redis-viewer/       Entrypoint (bubbletea TUI with syntax highlighting)
internal/
  cmd/                  Command types (RedisCmdData, Options)
  processor/            Input parsing, command building, dispatch
  redis/
    protocol.go         Adapter interface (AdapterRedisProtocol)
    go-redis/           go-redis/v9 adapter implementation
```

## Development

```bash
# run all tests
go test ./...

# run tests for a specific package
go test -v ./internal/processor/

# build
go build ./cmd/redis-viewer/
```

## License

MIT
