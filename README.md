# SocketX

A lightweight TCP proxy framework written in Go, designed to bypass network restrictions.

![CI](https://github.com/kissme666/socket-x/actions/workflows/ci.yml/badge.svg)

## Features

- TCP server / client mode
- YAML / JSON configuration
- Structured logging with rotation (slog + lumberjack)
- Cross-platform static binary
- Docker support

## Installation

**Binary**

Download the latest release from [GitHub Releases](https://github.com/kissme666/socket-x/releases).

**Docker**

```bash
docker pull ghcr.io/kissme666/socket-x:latest
```

**Build from source**

```bash
git clone https://github.com/kissme666/socket-x.git
cd socket-x
make build
```

## Usage

```bash
# Start server
./socketx server -c config.yaml

# Start client
./socketx client -c config.yaml
```

**Docker**

```bash
docker run -p 8888:8888 ghcr.io/kissme666/socket-x:latest

# Mount custom config
docker run -p 8888:8888 -v ./config.yaml:/config.yaml ghcr.io/kissme666/socket-x:latest
```

## Configuration

```yaml
server:
  addr: ":8888"
  max_conn: 1000

log:
  filename: "logs/socketx.log"
  level: "info"       # debug / info / warn / error
  format: "json"      # json / text
  max_size: 100       # MB per file
  max_backups: 7      # number of files to retain
  max_age: 30         # days to retain
  compress: true
```

## License

MIT
