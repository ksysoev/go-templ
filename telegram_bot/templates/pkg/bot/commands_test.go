package bot

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleCommand_Start(t *testing.T) {
	msg := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "/start",
		Entities: []tgbotapi.MessageEntity{
			{Type: "bot_command", Offset: 0, Length: 6},
		},
	}

	resp, err := handleCommand(t.Context(), msg)

	require.NoError(t, err)
	assert.Equal(t, int64(123), resp.ChatID)
	assert.NotEmpty(t, resp.Text)
}

func TestHandleCommand_Help(t *testing.T) {
	msg := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "/help",
		Entities: []tgbotapi.MessageEntity{
			{Type: "bot_command", Offset: 0, Length: 5},
		},
	}

	resp, err := handleCommand(t.Context(), msg)

	require.NoError(t, err)
	assert.Equal(t, int64(123), resp.ChatID)
	assert.Contains(t, resp.Text, "/start")
	assert.Contains(t, resp.Text, "/help")
}

func TestHandleCommand_Unknown(t *testing.T) {
	msg := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "/unknown",
		Entities: []tgbotapi.MessageEntity{
			{Type: "bot_command", Offset: 0, Length: 8},
		},
	}

	resp, err := handleCommand(t.Context(), msg)

	require.NoError(t, err)
	assert.Contains(t, resp.Text, "Unknown command")
}
