package middleware

import "net/http"

func ApplyMiddlewares(r http.Handler) http.Handler {
	r = SecurityHeadersMiddleware(r)

	return r
}
