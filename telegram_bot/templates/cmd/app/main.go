package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"{{ .Values.repo }}/pkg/cmd"
)

// version is the version of the application. It should be set at build time.
var (
	version = "dev"
	name    = "{{ .Values.command }}"
)

func main() {
	os.Exit(runApp())
}

func runApp() int {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	rootCmd := cmd.InitCommand(cmd.BuildInfo{
		Version: version,
		AppName: name,
	})

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		slog.Error("failed to execute command", slog.Any("error", err))

		return 1
	}

	return 0
}
