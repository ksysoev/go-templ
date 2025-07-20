package cmd

import (
	"context"
	"fmt"
)

func RunCommand(ctx context.Context, flags *cmdFlags) error {
	if err := initLogger(flags); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	_, err := loadConfig(flags)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}
