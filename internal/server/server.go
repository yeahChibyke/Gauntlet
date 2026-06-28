package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
	"github.com/yeahChibyke/Gauntlet/internal/service"
)

// NewHTTPServer creates and configures the Gauntlet HTTP server.
func NewHTTPServer(
	addr string,
	logger *slog.Logger,
	responseService *service.ResponseService,
) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var req responses.Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

			logger.Error(
				"failed to decode request",
				"error", err,
			)

			writeError(
				w,
				http.StatusBadRequest,
				"invalid_request",
				"failed to decode request body",
			)

			return
		}

		resp, err := responseService.Handle(
			r.Context(),
			&req,
		)
		if err != nil {

			logger.Error(
				"request failed",
				"error", err,
			)

			writeError(
				w,
				http.StatusBadRequest,
				"provider_error",
				err.Error(),
			)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {

			logger.Error(
				"failed to encode response",
				"error", err,
			)
		}
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

func writeError(
	w http.ResponseWriter,
	status int,
	errType string,
	message string,
) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]any{
		"error": map[string]any{
			"type":    errType,
			"message": message,
		},
	})
}
