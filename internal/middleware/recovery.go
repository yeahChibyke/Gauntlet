package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery prevents panics from crashing the HTTP server.
func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if recovered := recover(); recovered != nil {

					requestID := RequestIDFromContext(r.Context())

					logger.Error(
						"http handler panic",
						"request_id", requestID,
						"panic", recovered,
						"stack", string(debug.Stack()),
					)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)

					_ = json.NewEncoder(w).Encode(map[string]any{
						"error": map[string]any{
							"type":    "internal_error",
							"message": "internal server error",
						},
					})
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
