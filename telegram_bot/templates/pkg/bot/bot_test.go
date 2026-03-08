package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewService_NilConfig(t *testing.T) {
	svc := NewMockCoreService(t)

	_, err := NewService(nil, svc)

	assert.ErrorContains(t, err, "config cannot be nil")
}

func TestNewService_EmptyToken(t *testing.T) {
	svc := NewMockCoreService(t)

	_, err := NewService(&Config{TelegramToken: ""}, svc)

	assert.ErrorContains(t, err, "telegram token cannot be empty")
}

func TestNewService_NilCoreService(t *testing.T) {
	_, err := NewService(&Config{TelegramToken: "token"}, nil)

	assert.ErrorContains(t, err, "core service cannot be nil")
}

func TestNewService_InvalidToken(t *testing.T) {
	svc := NewMockCoreService(t)

	_, err := NewService(&Config{TelegramToken: "invalid-token"}, svc)

	require.Error(t, err)
	assert.ErrorContains(t, err, "failed to create Telegram bot")
}
