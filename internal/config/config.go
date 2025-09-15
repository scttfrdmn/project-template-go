// Package config handles application configuration.
package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration.
type Config struct {
	// Server configuration
	Port int    `json:"port"`
	Host string `json:"host"`

	// Application configuration
	Environment string `json:"environment"`
	LogLevel    string `json:"log_level"`
	Debug       bool   `json:"debug"`
}

// Load loads configuration from environment variables with sensible defaults.
func Load() (*Config, error) {
	cfg := &Config{
		Port:        8080,
		Host:        "0.0.0.0",
		Environment: "development",
		LogLevel:    "info",
		Debug:       false,
	}

	// Override with environment variables
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Port = p
		}
	}

	if host := os.Getenv("HOST"); host != "" {
		cfg.Host = host
	}

	if env := os.Getenv("ENVIRONMENT"); env != "" {
		cfg.Environment = env
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		cfg.LogLevel = logLevel
	}

	if debug := os.Getenv("DEBUG"); debug != "" {
		cfg.Debug = debug == "true"
	}

	return cfg, nil
}