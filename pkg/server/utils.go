package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RespondAggregateError(w http.ResponseWriter, statusCode int, err errors.AggregatedError) error {
	results := make([]ErrorResponse, len(err))
	for i := range err {
		switch e := err[i].(type) {
		case errors.AppError:
			results[i].Code = e.Code
			results[i].Field = e.Field
			results[i].Message = e.Message
		default:
			results[i].Code = errors.ErrorCodeInternalServerError
		}
	}

	return RespondJSON(w, statusCode, ErrorResponseWrapper{
		Errors: results,
	})
}

func RespondError(w http.ResponseWriter, err errors.AppError) error {
	return RespondAggregateError(w, err.StatusCode, errors.NewAggregatedError(err))
}

const bearer = "bearer "

func ExtractHeaderBearerToken(r *http.Request, header string) (string, bool) {
	headerToken, ok := r.Header[header]
	if !ok || len(headerToken) == 0 || headerToken[0] == "" {
		return "", false
	}

	token := headerToken[0]

	if strings.HasPrefix(strings.ToLower(token), bearer) {
		token = token[len(bearer):]
	}

	return token, true
}
