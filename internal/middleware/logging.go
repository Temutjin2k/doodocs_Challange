package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware logs HTTP requests and responses.
func LoggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log incoming request details
		logger.Info("Request received", "method", r.Method, "path", r.URL.Path, "ip", r.RemoteAddr)

		// Capture the response status code by wrapping the ResponseWriter
		responseWriter := &LoggingResponseWriter{ResponseWriter: w}
		next.ServeHTTP(responseWriter, r)

		// Log request processing time and response status
		logger.Info("Request processed", "method", r.Method, "path", r.URL.Path, "status", responseWriter.status, "duration", time.Since(start))
	})
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	status int
}

// Override
func (w *LoggingResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
