package server

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// NewHTTPServer creates and configures the Gauntlet HTTP server.
func NewHTTPServer(addr string, logger *slog.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error(
				"failed to read request body",
				"error", err,
			)

			http.Error(
				w,
				"failed to read request body",
				http.StatusInternalServerError,
			)

			return
		}
		defer r.Body.Close()

		logger.Info(
			"incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"remote_addr", r.RemoteAddr,
			"content_type", r.Header.Get("Content-Type"),
			"authorization_present", r.Header.Get("Authorization") != "",
			"user_agent", r.Header.Get("User-Agent"),
			"body", string(body),
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)

		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": map[string]any{
				"message": "Gauntlet translator not implemented",
				"type":    "not_implemented",
			},
		})
	})

	return &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}