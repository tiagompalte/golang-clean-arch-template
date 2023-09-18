package middleware

import (
	contextNative "context"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/constant"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/auth"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func ValidateExtractUserTokenMiddleware(header string, auth auth.Auth, findUserUUIDUseCase usecase.FindUserUUIDUseCase) server.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := server.ExtractHeaderBearerToken(r, header)
			if !ok {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			ctx := r.Context()

			content, isValid, err := auth.ValidateExtractToken(ctx, token)
			if err != nil || !isValid {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			user, err := findUserUUIDUseCase.Execute(ctx, content["user_id"].(string))
			if err != nil {
				server.RespondError(w, errors.NewAppUnauthorizedError())
				return
			}

			ctx = contextNative.WithValue(ctx, constant.ContextUser, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
