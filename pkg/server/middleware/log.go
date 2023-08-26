package middleware

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/logger"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func LogMiddleware(log logger.Logger) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Infof("%s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}
