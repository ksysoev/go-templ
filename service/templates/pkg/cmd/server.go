package cmd

import (
	"context"
	"fmt"
)

func RunCommand(ctx context.Context, args *cmdArgs) error {
	if err := initLogger(args); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	cfg, err := loadConfig(args)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return nil
}
