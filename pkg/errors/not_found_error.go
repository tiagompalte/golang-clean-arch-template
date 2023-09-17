package errors

import "net/http"

// ErrorCodeNotFound means the request resource was not found
const ErrorCodeNotFound = "not-found"

func NewAppNotFoundError(field string) AppError {
	return AppError{
		StatusCode: http.StatusNotFound,
		Code:       ErrorCodeNotFound,
		Field:      field,
	}
}
