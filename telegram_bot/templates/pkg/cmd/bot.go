package cmd

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"{{ .Values.repo }}/pkg/bot"
	"{{ .Values.repo }}/pkg/core"
	"{{ .Values.repo }}/pkg/prov/someapi"
	chat "{{ .Values.repo }}/pkg/repo/chat"
)

// RunBot initializes and runs the Telegram bot with the provided configuration.
// It sets up Redis, external provider, core service, and bot service, then starts the bot.
func RunBot(ctx context.Context, flags *cmdFlags) error {
	if err := initLogger(flags); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	cfg, err := loadConfig(flags)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	slog.InfoContext(ctx, "Starting {{ .Values.name }}", slog.String("version", flags.version))

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
	})

	defer func() {
		if err := rdb.Close(); err != nil {
			slog.Error("failed to close Redis connection", slog.Any("error", err))
		}
	}()

	someAPI := someapi.New(cfg.Provider.SomeAPI)
	chatRepo := chat.New(rdb)
	svc := core.New(chatRepo, someAPI)

	botSvc, err := bot.NewService(&cfg.Bot, svc)
	if err != nil {
		return fmt.Errorf("failed to create bot service: %w", err)
	}

	if err := botSvc.Run(ctx); err != nil {
		return fmt.Errorf("failed to run bot service: %w", err)
	}

	return nil
}
