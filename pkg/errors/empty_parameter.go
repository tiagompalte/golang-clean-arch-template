package errors

import (
	"net/http"
)

// ErrorCodeEmptyParameter means that parameter is empty
const ErrorCodeEmptyParameter = "empty-parameter"

func NewEmptyParameterError(field string) AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeEmptyParameter,
		Field:      field,
	}
}
