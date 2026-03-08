package middleware

import (
	"context"
	"errors"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ErrNilMessage is returned when a nil message is passed to a handler.
var ErrNilMessage = errors.New("message is nil")

// WithErrorHandling adds error handling middleware to a Handler.
// It intercepts errors returned by the next Handler and generates an appropriate error message response for the user.
// Context cancellation and deadline errors are passed through unchanged so callers can handle them explicitly.
// Returns a Middleware wrapping the original Handler with error handling logic.
func WithErrorHandling() Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(ctx context.Context, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
			if message == nil {
				return tgbotapi.MessageConfig{}, ErrNilMessage
			}

			msgConfig, err := next.Handle(ctx, message)
			if err != nil {
				// Pass context errors through so callers can distinguish cancellation from handler failures.
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return tgbotapi.MessageConfig{}, err
				}

				var chatID int64
				if message.Chat != nil {
					chatID = message.Chat.ID
				}

				slog.ErrorContext(ctx, "Failed to handle message", slog.Any("error", err))

				return tgbotapi.NewMessage(chatID, "Sorry, I encountered an error while processing your request. Please try again later."), nil
			}

			return msgConfig, nil
		})
	}
}
