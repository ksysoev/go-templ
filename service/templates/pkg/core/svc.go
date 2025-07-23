package core

import (
	"context"
	"fmt"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) CheckHealth(_ context.Context) error {
	return fmt.Errorf("not implemented")
}
