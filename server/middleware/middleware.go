package middleware

import (
	"net/http"
	"quiz-app/server/logger"
)

func ApplyMiddlewares(r http.Handler, log logger.Logger) http.Handler {
	r = SecurityHeadersMiddleware(r)
	r = GzipMiddleware(r)
	r = LoggingMiddleware(r, log)
	return r
}
