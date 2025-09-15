// Package handler contains the HTTP handlers and routing logic.
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/scttfrdmn/project-name/internal/service"
)

// Handler represents the HTTP handler layer.
type Handler struct {
	service *service.Service
	server  *http.Server
}

// New creates a new handler instance.
func New(svc *service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// Start starts the HTTP server.
func (h *Handler) Start(ctx context.Context) error {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", h.handleHealth)
	mux.HandleFunc("/info", h.handleInfo)
	mux.HandleFunc("/api/process", h.handleProcess)

	h.server = &http.Server{
		Addr:         ":8080", // Should come from config
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Printf("Server starting on %s\n", h.server.Addr)
	return h.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (h *Handler) Shutdown(ctx context.Context) error {
	if h.server != nil {
		return h.server.Shutdown(ctx)
	}
	return nil
}

// handleHealth handles health check requests.
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status, err := h.service.Health(r.Context())
	if err != nil {
		http.Error(w, "Service unhealthy", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": status})
}

// handleInfo handles info requests.
func (h *Handler) handleInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	info, err := h.service.GetInfo(r.Context())
	if err != nil {
		http.Error(w, "Failed to get service info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// handleProcess handles data processing requests.
func (h *Handler) handleProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := h.service.ProcessData(r.Context(), req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}