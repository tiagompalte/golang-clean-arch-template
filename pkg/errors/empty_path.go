package errors

import (
	"net/http"
)

// ErrorCodeEmptyPath means that path is empty
const ErrorCodeEmptyPath = "empty-path"

func NewEmptyPathError(field string) AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeEmptyPath,
		Field:      field,
	}
}
