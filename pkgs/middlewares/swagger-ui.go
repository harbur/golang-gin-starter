package middlewares

import (
	"net/http"
	"strings"
)

// SwaggerUIMiddleware provides swagger-ui page on /
func SwaggerUIMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") || r.URL.Path == "/swagger.json" || r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("./swagger-ui/")).ServeHTTP(w, r)
		}
	})
}
