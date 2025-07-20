package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type BuildInfo struct {
	Version string
}
type cmdFlags struct {
	Version    string
	ConfigPath string `mapstructure:"config"`
	LogLevel   string `mapstructure:"log_level"`
	TextFormat bool   `mapstructure:"log_text"`
}

// InitCommand initializes the root command of the CLI application with its subcommands and flags.
func InitCommand(build BuildInfo) cobra.Command {
	arg := cmdFlags{
		Version: build.Version,
	}

	cmd := cobra.Command{
		Use:   "cli",
		Short: "",
		Long:  "",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return RunCommand(cmd, &arg)
		},
	}

	cmd.PersistentFlags().StringVar(&arg.LogLevel, "log-level", "info", "log level (debug, info, warn, error)")
	cmd.PersistentFlags().BoolVar(&arg.TextFormat, "log-text", true, "log in text format, otherwise JSON")

	for _, name := range []string{"log_level", "log_text"} {
		if err := viper.BindEnv(name); err != nil {
			slog.Error("failed to bind env var", "name", name, "error", err)
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&arg); err != nil {
		slog.Error("failed to unmarshal env vars", "error", err)
	}

	return cmd
}
