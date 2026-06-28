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
func NewHTTPServer(addr string, logger *slog.Logger) *http.Server {
	responseService := service.NewResponseService()

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

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

		canonicalReq, err := responseService.Handle(&req)
		if err != nil {
			logger.Error(
				"failed to translate request",
				"error", err,
			)

			writeError(
				w,
				http.StatusBadRequest,
				"translation_error",
				err.Error(),
			)
			return
		}

		logger.Info(
			"canonical request",
			"request", canonicalReq,
		)

		writeError(
			w,
			http.StatusNotImplemented,
			"not_implemented",
			"Gauntlet translator not implemented",
		)
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
