// Package main is the entry point for the application.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/scttfrdmn/project-name/internal/config"
	"github.com/scttfrdmn/project-name/internal/handler"
	"github.com/scttfrdmn/project-name/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize service
	svc := service.New(cfg)

	// Initialize handler
	h := handler.New(svc)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the application
	go func() {
		if err := h.Start(ctx); err != nil {
			log.Printf("Application error: %v", err)
			cancel()
		}
	}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		fmt.Println("\nReceived interrupt signal, shutting down gracefully...")
	case <-ctx.Done():
		fmt.Println("Application context cancelled, shutting down...")
	}

	// Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := h.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	fmt.Println("Application stopped")
}