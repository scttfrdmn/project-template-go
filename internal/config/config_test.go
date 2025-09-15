package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default values
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.Port != 8080 {
		t.Fatalf("Expected default port to be 8080, got %d", cfg.Port)
	}

	if cfg.Host != "0.0.0.0" {
		t.Fatalf("Expected default host to be '0.0.0.0', got %s", cfg.Host)
	}

	if cfg.Environment != "development" {
		t.Fatalf("Expected default environment to be 'development', got %s", cfg.Environment)
	}

	if cfg.LogLevel != "info" {
		t.Fatalf("Expected default log level to be 'info', got %s", cfg.LogLevel)
	}

	if cfg.Debug != false {
		t.Fatalf("Expected default debug to be false, got %t", cfg.Debug)
	}
}

func TestLoadWithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("HOST", "localhost")
	os.Setenv("ENVIRONMENT", "test")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("DEBUG", "true")

	// Clean up after test
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("HOST")
		os.Unsetenv("ENVIRONMENT")
		os.Unsetenv("LOG_LEVEL")
		os.Unsetenv("DEBUG")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if cfg.Port != 9090 {
		t.Fatalf("Expected port to be 9090, got %d", cfg.Port)
	}

	if cfg.Host != "localhost" {
		t.Fatalf("Expected host to be 'localhost', got %s", cfg.Host)
	}

	if cfg.Environment != "test" {
		t.Fatalf("Expected environment to be 'test', got %s", cfg.Environment)
	}

	if cfg.LogLevel != "debug" {
		t.Fatalf("Expected log level to be 'debug', got %s", cfg.LogLevel)
	}

	if cfg.Debug != true {
		t.Fatalf("Expected debug to be true, got %t", cfg.Debug)
	}
}

func TestLoadWithInvalidPort(t *testing.T) {
	// Set invalid port
	os.Setenv("PORT", "invalid")
	defer os.Unsetenv("PORT")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Should fall back to default port
	if cfg.Port != 8080 {
		t.Fatalf("Expected port to fall back to 8080, got %d", cfg.Port)
	}
}