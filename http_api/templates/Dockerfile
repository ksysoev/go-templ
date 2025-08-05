FROM golang:1.24.4-alpine AS builder

ARG VERSION=${VERSION}

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o {{ .Values.command }} -ldflags "-X main.version=$VERSION" ./cmd/{{ .Values.command }}/main.go

FROM scratch

COPY --from=builder /app/{{ .Values.command }} .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/{{ .Values.command }}"]
