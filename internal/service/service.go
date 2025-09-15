// Package service contains the business logic layer.
package service

import (
	"context"
	"fmt"

	"github.com/scttfrdmn/project-name/internal/config"
)

// Service represents the application service layer.
type Service struct {
	config *config.Config
}

// New creates a new service instance.
func New(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
	}
}

// Health returns the health status of the service.
func (s *Service) Health(ctx context.Context) (string, error) {
	return "healthy", nil
}

// GetInfo returns basic information about the service.
func (s *Service) GetInfo(ctx context.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"service":     "project-name",
		"environment": s.config.Environment,
		"version":     "1.0.0", // This should come from build info
	}, nil
}

// Example business logic method
func (s *Service) ProcessData(ctx context.Context, data string) (string, error) {
	if data == "" {
		return "", fmt.Errorf("data cannot be empty")
	}

	// Add your business logic here
	processed := fmt.Sprintf("processed: %s", data)
	return processed, nil
}