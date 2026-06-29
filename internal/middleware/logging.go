package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Logging logs every HTTP request after it has completed.
func Logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			recorder := NewResponseRecorder(w)

			start := time.Now()

			next.ServeHTTP(recorder, r)

			requestID := RequestIDFromContext(r.Context())

			logger.Info(
				"http request completed",
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
				"status", recorder.StatusCode,
				"bytes", recorder.Bytes,
				"duration_ms", time.Since(start).Milliseconds(),
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		})
	}
}
