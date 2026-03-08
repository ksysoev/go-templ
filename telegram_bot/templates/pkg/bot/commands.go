package bot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handleCommand dispatches incoming bot commands to the appropriate handler.
// Unknown commands return a friendly error message to the user.
func handleCommand(ctx context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	switch msg.Command() {
	case "start":
		return handleStart(ctx, msg)
	case "help":
		return handleHelp(ctx, msg)
	default:
		return tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Unknown command: /%s", msg.Command())), nil
	}
}

// handleStart sends the welcome message when the user initiates a conversation.
func handleStart(_ context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := `Welcome! I'm your bot assistant.

Send me a message and I'll respond. Use /help to see available commands.`

	return tgbotapi.NewMessage(msg.Chat.ID, text), nil
}

// handleHelp sends the help message listing available commands.
func handleHelp(_ context.Context, msg *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	text := `<b>Available Commands:</b>
/start - Start the conversation
/help - Show this help message`

	resp := tgbotapi.NewMessage(msg.Chat.ID, text)
	resp.ParseMode = "HTML"

	return resp, nil
}
