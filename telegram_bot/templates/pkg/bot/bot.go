// Package bot provides the Telegram bot service implementation.
package bot

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"{{ .Values.repo }}/pkg/bot/middleware"
)

const (
	requestTimeout = 120 * time.Second
	typingInterval = 5 * time.Second
	updateTimeout  = 30
)

// BotAPI defines the Telegram bot API capabilities used by the service.
type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	StopReceivingUpdates()
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

// Service defines the interface for the bot service lifecycle.
type Service interface {
	Run(ctx context.Context) error
}

// Config holds the configuration for the Telegram bot.
type Config struct {
	TelegramToken string `mapstructure:"telegram_token"`
}

// ServiceImpl is the main Telegram bot service.
type ServiceImpl struct {
	bot     BotAPI
	handler middleware.Handler
}

// NewService creates a new bot service with the given configuration and core service.
// It returns an error if the configuration is invalid or the bot API connection fails.
func NewService(cfg *Config, svc CoreService) (*ServiceImpl, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	if cfg.TelegramToken == "" {
		return nil, fmt.Errorf("telegram token cannot be empty")
	}

	if svc == nil {
		return nil, fmt.Errorf("core service cannot be nil")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create Telegram bot: %w", err)
	}

	s := &ServiceImpl{
		bot: bot,
	}

	s.handler = s.setupHandler(svc)

	return s, nil
}

// Run starts the bot's update polling loop and blocks until the context is cancelled.
// It handles graceful shutdown by waiting for in-flight requests to complete.
func (s *ServiceImpl) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting Telegram bot")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = updateTimeout

	updates := s.bot.GetUpdatesChan(updateConfig)

	var wg sync.WaitGroup

	for {
		select {
		case update, ok := <-updates:
			if !ok {
				return nil
			}

			wg.Add(1)

			go func() {
				defer wg.Done()

				reqCtx, cancel := context.WithTimeout(ctx, requestTimeout)

				//nolint:staticcheck // standard pattern for request-scoped values
				reqCtx = context.WithValue(reqCtx, "req_id", uuid.New().String())

				defer cancel()

				s.processUpdate(reqCtx, &update)
			}()

		case <-ctx.Done():
			slog.Info("Starting graceful shutdown")
			s.bot.StopReceivingUpdates()

			done := make(chan struct{})

			go func() {
				wg.Wait()
				close(done)
			}()

			select {
			case <-done:
				slog.InfoContext(ctx, "Graceful shutdown completed")
			case <-time.After(requestTimeout):
				slog.Warn("Graceful shutdown timed out")
			}

			return nil
		}
	}
}

// processUpdate handles a single Telegram update, dispatching it to the handler
// and sending the response back to the user.
func (s *ServiceImpl) processUpdate(ctx context.Context, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	msg := update.Message

	//nolint:staticcheck // standard pattern for request-scoped values
	ctx = context.WithValue(ctx, "chat_id", fmt.Sprintf("%d", msg.Chat.ID))

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg.Add(1)

	go func() {
		defer wg.Done()

		s.keepTyping(ctx, msg.Chat.ID, typingInterval)
	}()

	defer wg.Wait()

	msgConfig, err := s.handler.Handle(ctx, msg)

	if errors.Is(err, context.Canceled) {
		slog.InfoContext(ctx, "Request cancelled",
			slog.Int64("chat_id", msg.Chat.ID),
		)

		return
	} else if err != nil {
		slog.ErrorContext(ctx, "Unexpected error processing update",
			slog.Any("error", err),
		)

		return
	}

	if msgConfig.Text == "" {
		return
	}

	cancel()

	if _, err := s.bot.Send(msgConfig); err != nil {
		slog.ErrorContext(ctx, "Failed to send message",
			slog.Any("error", err),
		)
	}
}

// sendTyping sends a "typing" action to the specified chat.
func (s *ServiceImpl) sendTyping(ctx context.Context, chatID int64) {
	typing := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	if _, err := s.bot.Request(typing); err != nil {
		slog.ErrorContext(ctx, "Failed to send typing action",
			slog.Any("error", err),
		)
	}
}

// keepTyping continuously sends typing notifications at the given interval until the context is cancelled.
func (s *ServiceImpl) keepTyping(ctx context.Context, chatID int64, interval time.Duration) {
	s.sendTyping(ctx, chatID)

	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			s.sendTyping(ctx, chatID)
		}
	}
}
