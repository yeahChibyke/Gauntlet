package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/yeahChibyke/Gauntlet/internal/middleware"
	"github.com/yeahChibyke/Gauntlet/internal/protocol/responses"
	"github.com/yeahChibyke/Gauntlet/internal/provider"
	"github.com/yeahChibyke/Gauntlet/internal/service"
	"github.com/yeahChibyke/Gauntlet/internal/translate"
)

// NewHTTPServer creates and configures the Gauntlet HTTP server.
func NewHTTPServer(
	addr string,
	logger *slog.Logger,
	responseService *service.ResponseService,
) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/models", handleModels)

	mux.HandleFunc("/v1/responses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		requestID := middleware.RequestIDFromContext(r.Context())

		reqLogger := logger.With(
			"request_id", requestID,
		)

		var req responses.Request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			reqLogger.Error(
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

		if req.Stream {

			stream, err := responseService.HandleStream(
				r.Context(),
				&req,
			)
			if err != nil {

				reqLogger.Error(
					"stream request failed",
					"error", err,
				)

				var providerErr *provider.Error

				if errors.As(err, &providerErr) {
					writeError(
						w,
						providerErr.Status,
						"provider_error",
						providerErr.Message,
					)
					return
				}

				writeError(
					w,
					http.StatusInternalServerError,
					"internal_error",
					"internal server error",
				)

				return
			}

			defer stream.Close()

			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			if err := writeResponseCreated(w); err != nil {
				return
			}

			for {

				canonicalResp, ok, err := stream.Next(r.Context())
				if err != nil {
					reqLogger.Error(
						"stream failed",
						"error", err,
					)
					return
				}

				if !ok {
					break
				}

				if err := writeResponseDelta(
					w,
					translate.ToResponses(canonicalResp),
				); err != nil {
					return
				}
			}

			_ = writeStreamDone(w)

			return
		}

		resp, err := responseService.Handle(r.Context(), &req)
		if err != nil {

			reqLogger.Error(
				"request failed",
				"error", err,
			)

			var providerErr *provider.Error

			if errors.As(err, &providerErr) {

				writeError(
					w,
					providerErr.Status,
					"provider_error",
					providerErr.Message,
				)

				return
			}

			writeError(
				w,
				http.StatusInternalServerError,
				"internal_error",
				"internal server error",
			)

			return
		}

		reqLogger.Info(
			"provider request completed",
			"model", req.Model,
		)

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			reqLogger.Error(
				"failed to encode response",
				"error", err,
			)
		}
	})

	// Compose middleware from inside to outside:
	//
	// RequestID
	//      ↓
	// Recovery
	//      ↓
	// Logging
	//      ↓
	// HTTP Server
	handler := middleware.Logging(logger)(
		middleware.Recovery(logger)(
			middleware.RequestID(mux),
		),
	)

	return &http.Server{
		Addr:              addr,
		Handler:           handler,
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
