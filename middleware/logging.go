package middleware

import (
	"bytes"
	"net/http"
	"quiz-app/logger"
	"time"
)

func LoggingMiddleware(next http.Handler, log logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Info("Incoming request", map[string]interface{}{
			"method":  r.Method,
			"path":    r.URL.Path,
			"headers": r.Header,
		})

		ww := &responseWriterWrapper{ResponseWriter: w}

		// Serve the request
		next.ServeHTTP(ww, r)

		// Log the outgoing response
		duration := time.Since(start)
		log.Info("Request processed", map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   ww.statusCode,
			"duration": duration,
		})
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (ww *responseWriterWrapper) WriteHeader(code int) {
	ww.statusCode = code
	ww.ResponseWriter.WriteHeader(code)
}

func (ww *responseWriterWrapper) Write(p []byte) (int, error) {
	ww.body.Write(p)
	return ww.ResponseWriter.Write(p)
}
