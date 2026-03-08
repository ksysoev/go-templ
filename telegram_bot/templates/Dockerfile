FROM golang:1.24.4-alpine AS builder

ARG VERSION=${VERSION}

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o {{ .Values.command }} -ldflags "-X main.version=$VERSION -X main.name={{ .Values.command }}" ./cmd/{{ .Values.command }}/main.go

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/{{ .Values.command }} .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER appuser

ENTRYPOINT ["/{{ .Values.command }}"]
