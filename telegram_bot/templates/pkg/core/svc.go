// Package core provides core service logic and interfaces.
package core

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// chatRepo defines the interface for chat repository operations.
type chatRepo interface {
	CheckHealth(ctx context.Context) error
}

// someAPIProv defines the interface for a provider that can check health status.
type someAPIProv interface {
	CheckHealth(ctx context.Context) error
}

// Service encapsulates core business logic and dependencies.
type Service struct {
	chats   chatRepo
	someAPI someAPIProv
}

// New creates a new Service instance with the provided chatRepo and someAPI provider.
func New(chats chatRepo, someAPI someAPIProv) *Service {
	return &Service{
		chats:   chats,
		someAPI: someAPI,
	}
}

// CheckHealth checks the health of the core service and its dependencies.
func (s *Service) CheckHealth(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error { return s.someAPI.CheckHealth(ctx) })
	eg.Go(func() error { return s.chats.CheckHealth(ctx) })

	return eg.Wait()
}
