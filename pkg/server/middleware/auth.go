package middleware

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func AuthMiddleware(header string, auth auth.Auth) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := r.Header[header]
			if !ok || len(token) == 0 || token[0] == "" {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			isValid, err := auth.ValidateToken(r.Context(), token[0])
			if err != nil || !isValid {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
