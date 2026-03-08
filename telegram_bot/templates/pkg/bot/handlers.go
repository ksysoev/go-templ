package bot

import (
	"context"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"{{ .Values.repo }}/pkg/bot/middleware"
)

// CoreService defines the interface for the core business logic the bot depends on.
type CoreService interface {
	CheckHealth(ctx context.Context) error
}

// setupHandler initializes and configures the request handler with middleware.
// It applies throttling, metrics collection, and error handling middleware.
// Returns a Handler that processes messages with the applied middleware stack.
func (s *ServiceImpl) setupHandler(svc CoreService) middleware.Handler {
	h := middleware.Use(
		newBaseHandler(svc),
		middleware.WithThrottler(30),
		middleware.WithMetrics(),
		middleware.WithErrorHandling(),
	)

	return h
}

// baseHandler wraps the core service and handles message dispatch.
type baseHandler struct {
	svc CoreService
}

func newBaseHandler(svc CoreService) middleware.Handler {
	return &baseHandler{svc: svc}
}

// Handle processes an incoming Telegram message and generates an appropriate response.
// It dispatches to command handling or text message handling based on the message content.
func (h *baseHandler) Handle(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	slog.DebugContext(ctx, "Received message", slog.String("text", msg.Text))

	if msg.From == nil {
		return tgbotapi.MessageConfig{}, fmt.Errorf("message has no sender")
	}

	if msg.Text == "" {
		return tgbotapi.MessageConfig{}, nil
	}

	if msg.Command() != "" {
		return handleCommand(ctx, msg)
	}

	// Placeholder: handle plain text messages here.
	// Replace this with your actual business logic using h.svc.
	return tgbotapi.NewMessage(msg.Chat.ID, msg.Text), nil
}
