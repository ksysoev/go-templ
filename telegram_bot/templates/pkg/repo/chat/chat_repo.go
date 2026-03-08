// Package chat provides repository implementations for chat-related data access.
package chat

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// chatDAO defines the interface for chat data access operations.
type chatDAO interface {
	Ping(ctx context.Context) *redis.StatusCmd
}

// ChatRepo provides methods to interact with the chat data store.
type ChatRepo struct {
	dao chatDAO
}

// New creates a new instance of ChatRepo using the provided chatDAO.
// It returns a pointer to the initialized ChatRepo.
func New(dao chatDAO) *ChatRepo {
	return &ChatRepo{
		dao: dao,
	}
}

// CheckHealth verifies the health of the chat repository by pinging the underlying DAO.
// It returns an error if the health check fails.
func (c *ChatRepo) CheckHealth(ctx context.Context) error {
	res := c.dao.Ping(ctx)

	if err := res.Err(); err != nil {
		return fmt.Errorf("fail to check health for chat repo: %w", err)
	}

	return nil
}
