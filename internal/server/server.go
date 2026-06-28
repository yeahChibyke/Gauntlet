package server

import (
	"io"
	"log/slog"
	"net/http"
	"time"
)

// NewHTTPServer creates and configures the Gauntlet HTTP server.
func NewHTTPServer(addr string, logger *slog.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(
			"incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"remote_addr", r.RemoteAddr,
		)

		logger.Info(
			"request headers",
			"headers", r.Header,
		)

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
			"request body",
			"body", string(body),
		)

		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusNotImplemented)

		_, _ = w.Write([]byte("not implemented\n"))
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