# {{ .Values.name }}

[![Tests](https://{{ .Values.repo }}/actions/workflows/tests.yml/badge.svg)](https://{{ .Values.repo }}/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/{{ .Values.repo }})](https://goreportcard.com/report/{{ .Values.repo }})
[![Go Reference](https://pkg.go.dev/badge/{{ .Values.repo }}.svg)](https://pkg.go.dev/{{ .Values.repo }})
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

{{ .Values.description }}

## Installation

### Using Go

If you have Go installed, you can install {{ .Values.name }} directly:

```sh
go install {{ .Values.repo }}/cmd/{{ .Values.command }}@latest
```

## Building from Source

```sh
CGO_ENABLED=0 go build -o {{ .Values.command }} -ldflags "-X main.version=dev -X main.name={{ .Values.command }}" ./cmd/{{ .Values.command }}/main.go
```

## Configuration

Copy `runtime/config.yml` and fill in your Telegram bot token:

```yaml
bot:
  telegram_token: your-telegram-bot-token

redis:
  addr: 127.0.0.1:6379
  password:

provider:
  some_api:
    base_url: http://example.com
```

## Using

```sh
{{ .Values.command }} --log-level=debug --log-text=true --config=runtime/config.yml
```

## License

{{ .Values.name }} is licensed under the MIT License. See the LICENSE file for more details.
