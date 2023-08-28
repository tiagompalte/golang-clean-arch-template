package server

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
	statusCode, ok := HttpStatusCode[err.Code]
	if !ok {
		statusCode = http.StatusInternalServerError
	}
	return RespondAggregateError(w, statusCode, errors.NewAggregatedError(err))
}
