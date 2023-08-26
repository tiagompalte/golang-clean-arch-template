package middleware

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func AuthMiddleware(header string) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth, ok := r.Header[header]
			// TODO: implements validation
			if !ok || len(auth) == 0 {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
