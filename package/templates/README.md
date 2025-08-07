# {{ .Values.name }}

[![Tests](https://{{ .Values.repo }}/actions/workflows/tests.yml/badge.svg)](https://{{ .Values.repo }}/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/{{ .Values.repo }})](https://goreportcard.com/report/{{ .Values.repo }})
[![Go Reference](https://pkg.go.dev/badge/{{ .Values.repo }}.svg)](https://pkg.go.dev/{{ .Values.repo }})
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

{{ .Values.description }}

## Installation

```sh
go get {{ .Values.repo }}@latest
```

## License

{{ .Values.name }} is licensed under the MIT License. See the LICENSE file for more details.
