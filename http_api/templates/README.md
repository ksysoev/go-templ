# {{ .Values.name }}

[![Tests](https://{{ .Values.repo }}/actions/workflows/tests.yml/badge.svg)](https://{{ .Values.repo }}/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/{{ .Values.name }})](https://goreportcard.com/report/{{ .Values.name }})
[![Go Reference](https://pkg.go.dev/badge/{{ .Values.name }}.svg)](https://pkg.go.dev/{{ .Values.name }})
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

{{ .Values.description }}

## Installation

### Using Go

If you have Go installed, you can install {{ .Values.name }} directly:

```bash
go install {{ .Values.repo }}/cmd/{{ .Values.command }}@latest
```

## License

{{ .Values.name }} is licensed under the MIT License. See the LICENSE file for more details.
