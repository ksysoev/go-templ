package cmd

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"{{ .Values.repo }}/pkg/api"
	"{{ .Values.repo }}/pkg/core"
	"{{ .Values.repo }}/pkg/prov/someapi"
	"{{ .Values.repo }}/pkg/repo/user"
)

// RunCommand initializes the logger, loads configuration, creates the core and API services,
// and starts the API service. It returns an error if any step fails.
func RunCommand(ctx context.Context, flags *cmdFlags) error {
	if err := initLogger(flags); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	cfg, err := loadConfig(flags)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
	})

	someAPI := someapi.New(cfg.Provider.SomeAPI)
	userRepo := user.New(rdb)

	svc := core.New(userRepo, someAPI)
	apiSvc, err := api.New(cfg.API, svc)
	if err != nil {
		return fmt.Errorf("failed to create API service: %w", err)
	}

	return apiSvc.Run(ctx)
}
