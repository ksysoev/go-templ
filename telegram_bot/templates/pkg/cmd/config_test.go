package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"{{ .Values.repo }}/pkg/bot"
	"{{ .Values.repo }}/pkg/prov/someapi"
)

func TestLoadConfig(t *testing.T) {
	const validConfig = `
bot:
  telegram_token: "test-token"
redis:
  addr: "localhost:6379"
  password: "testpassword"
provider:
  some_api:
    base_url: "https://api.example.com"
`

	tests := []struct {
		envVars      map[string]string
		expectConfig *appConfig
		name         string
		configData   string
		expectError  bool
	}{
		{
			name:        "valid config file",
			envVars:     nil,
			expectError: false,
			configData:  validConfig,
			expectConfig: &appConfig{
				Bot: bot.Config{
					TelegramToken: "test-token",
				},
				Redis: RedisConfig{
					Addr:     "localhost:6379",
					Password: "testpassword",
				},
				Provider: Provider{
					SomeAPI: someapi.Config{
						BaseURL: "https://api.example.com",
					},
				},
			},
		},
		{
			name:        "missing config file",
			envVars:     nil,
			expectError: true,
		},
		{
			name:        "unparseable config file",
			envVars:     nil,
			expectError: true,
			configData:  `invalid yaml`,
		},
		{
			name: "valid config with environment overrides",
			envVars: map[string]string{
				"BOT_TELEGRAM_TOKEN":         "env-token",
				"PROVIDER_SOME_API_BASE_URL": "https://test.com",
			},
			expectError: false,
			configData:  validConfig,
			expectConfig: &appConfig{
				Bot: bot.Config{
					TelegramToken: "env-token",
				},
				Redis: RedisConfig{
					Addr:     "localhost:6379",
					Password: "testpassword",
				},
				Provider: Provider{
					SomeAPI: someapi.Config{
						BaseURL: "https://test.com",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			configPath := filepath.Join(tmpDir, "config.yaml")

			if tt.configData != "" {
				err := os.WriteFile(configPath, []byte(tt.configData), 0o600)
				require.NoError(t, err)
			}

			// Set up environment variables
			if tt.envVars != nil {
				for key, value := range tt.envVars {
					t.Setenv(key, value)
				}
			}

			arg := &cmdFlags{
				ConfigPath: configPath,
			}

			cfg, err := loadConfig(arg)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, cfg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectConfig, cfg)
			}
		})
	}
}
