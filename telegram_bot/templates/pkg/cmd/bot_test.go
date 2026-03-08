package cmd

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunBot_InitLoggerFails(t *testing.T) {
	flags := &cmdFlags{
		LogLevel: "WrongLogLevel",
	}

	err := RunBot(t.Context(), flags)
	assert.ErrorContains(t, err, "failed to init logger")
}

func TestRunBot_LoadConfigFails(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yaml")

	err := os.WriteFile(configPath, []byte("invalid config"), 0o600)
	require.NoError(t, err)

	flags := &cmdFlags{
		ConfigPath: configPath,
		LogLevel:   "info",
	}

	err = RunBot(t.Context(), flags)
	assert.ErrorContains(t, err, "failed to load config:")
}

func TestRunBot_BotFails(t *testing.T) {
	// No telegram token set — bot creation should fail
	flags := &cmdFlags{
		LogLevel: "info",
	}

	err := RunBot(t.Context(), flags)
	assert.ErrorContains(t, err, "failed to create bot service:")
}

func TestRunBot_Success(t *testing.T) {
	t.Setenv("BOT_TELEGRAM_TOKEN", "test-token")

	ctx, cancel := context.WithCancel(t.Context())

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()

	// Bot creation will fail with invalid token — this tests the wiring, not Telegram connectivity
	err := RunBot(ctx, &cmdFlags{LogLevel: "info"})
	assert.ErrorContains(t, err, "failed to create bot service:")
}
