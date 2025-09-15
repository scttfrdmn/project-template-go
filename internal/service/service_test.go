package service

import (
	"context"
	"testing"

	"github.com/scttfrdmn/project-name/internal/config"
)

func TestNew(t *testing.T) {
	cfg := &config.Config{
		Port:        8080,
		Host:        "localhost",
		Environment: "test",
		LogLevel:    "info",
		Debug:       false,
	}

	svc := New(cfg)
	if svc == nil {
		t.Fatal("Expected service to be created, got nil")
	}

	if svc.config != cfg {
		t.Fatal("Expected service config to match provided config")
	}
}

func TestHealth(t *testing.T) {
	cfg := &config.Config{}
	svc := New(cfg)

	ctx := context.Background()
	status, err := svc.Health(ctx)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if status != "healthy" {
		t.Fatalf("Expected status to be 'healthy', got %s", status)
	}
}

func TestGetInfo(t *testing.T) {
	cfg := &config.Config{
		Environment: "test",
	}
	svc := New(cfg)

	ctx := context.Background()
	info, err := svc.GetInfo(ctx)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info["service"] != "project-name" {
		t.Fatalf("Expected service name to be 'project-name', got %v", info["service"])
	}

	if info["environment"] != "test" {
		t.Fatalf("Expected environment to be 'test', got %v", info["environment"])
	}

	if info["version"] != "1.0.0" {
		t.Fatalf("Expected version to be '1.0.0', got %v", info["version"])
	}
}

func TestProcessData(t *testing.T) {
	cfg := &config.Config{}
	svc := New(cfg)

	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "valid data",
			input:       "test data",
			expected:    "processed: test data",
			expectError: false,
		},
		{
			name:        "empty data",
			input:       "",
			expected:    "",
			expectError: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := svc.ProcessData(ctx, tt.input)

			if tt.expectError {
				if err == nil {
					t.Fatal("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if result != tt.expected {
				t.Fatalf("Expected result to be %s, got %s", tt.expected, result)
			}
		})
	}
}