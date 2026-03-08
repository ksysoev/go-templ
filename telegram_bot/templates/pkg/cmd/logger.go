package cmd

import (
	"context"
	"log/slog"
	"os"
)

// ContextHandler is a custom slog.Handler that enriches log records with application-specific attributes.
// It embeds a slog.Handler and adds attributes like application name and version, as well as request-specific context data.
type ContextHandler struct {
	slog.Handler
	ver string
	app string
}

// Handle processes a log record by enriching it with context and application-specific attributes.
// It adds attributes such as "req_id" and "chat_id" from the context, "app", and "ver" before delegating to the embedded handler.
// Returns error if the embedded handler fails.

//nolint:gocritic // ignore this linting rule
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value("req_id").(string); ok { //nolint:staticcheck // standard pattern for request-scoped values
		r.AddAttrs(slog.String("req_id", reqID))
	}

	if chatID, ok := ctx.Value("chat_id").(string); ok { //nolint:staticcheck // standard pattern for request-scoped values
		r.AddAttrs(slog.String("chat_id", chatID))
	}

	r.AddAttrs(slog.String("app", h.app), slog.String("ver", h.ver))

	return h.Handler.Handle(ctx, r)
}

// WithAttrs returns a new ContextHandler that wraps the result of the embedded handler's WithAttrs call.
// This ensures the context enrichment (req_id, chat_id, app, ver) is preserved after logger.With(...) calls.
func (h ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return ContextHandler{
		Handler: h.Handler.WithAttrs(attrs),
		ver:     h.ver,
		app:     h.app,
	}
}

// WithGroup returns a new ContextHandler that wraps the result of the embedded handler's WithGroup call.
// This ensures the context enrichment (req_id, chat_id, app, ver) is preserved after logger.WithGroup(...) calls.
func (h ContextHandler) WithGroup(name string) slog.Handler {
	return ContextHandler{
		Handler: h.Handler.WithGroup(name),
		ver:     h.ver,
		app:     h.app,
	}
}

// initLogger initializes the default logger for the application using slog.
// It returns an error if the log level string is invalid.
func initLogger(flags *cmdFlags) error {
	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(flags.LogLevel)); err != nil {
		return err
	}

	options := &slog.HandlerOptions{
		Level: logLevel,
	}

	var logHandler slog.Handler
	if flags.TextFormat {
		logHandler = slog.NewTextHandler(os.Stdout, options)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, options)
	}

	ctxHandler := &ContextHandler{
		Handler: logHandler,
		ver:     flags.version,
		app:     flags.appName,
	}

	logger := slog.New(ctxHandler)

	slog.SetDefault(logger)

	return nil
}
