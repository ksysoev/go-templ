package bot

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBaseHandler_Handle_NilFrom(t *testing.T) {
	svc := NewMockCoreService(t)
	h := newBaseHandler(svc)

	msg := &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 123},
	}

	_, err := h.Handle(t.Context(), msg)

	require.Error(t, err)
	assert.ErrorContains(t, err, "message has no sender")
}

func TestBaseHandler_Handle_EmptyText(t *testing.T) {
	svc := NewMockCoreService(t)
	h := newBaseHandler(svc)

	msg := &tgbotapi.Message{
		From: &tgbotapi.User{ID: 1},
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "",
	}

	resp, err := h.Handle(t.Context(), msg)

	require.NoError(t, err)
	assert.Empty(t, resp.Text)
}

func TestBaseHandler_Handle_PlainText(t *testing.T) {
	svc := NewMockCoreService(t)
	h := newBaseHandler(svc)

	msg := &tgbotapi.Message{
		From: &tgbotapi.User{ID: 1},
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "hello",
	}

	resp, err := h.Handle(t.Context(), msg)

	require.NoError(t, err)
	assert.Equal(t, "hello", resp.Text)
}

func TestBaseHandler_Handle_Command(t *testing.T) {
	svc := NewMockCoreService(t)
	h := newBaseHandler(svc)

	msg := &tgbotapi.Message{
		From: &tgbotapi.User{ID: 1},
		Chat: &tgbotapi.Chat{ID: 123},
		Text: "/start",
		Entities: []tgbotapi.MessageEntity{
			{Type: "bot_command", Offset: 0, Length: 6},
		},
	}

	resp, err := h.Handle(t.Context(), msg)

	require.NoError(t, err)
	assert.NotEmpty(t, resp.Text)
}
