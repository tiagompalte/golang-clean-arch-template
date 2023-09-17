package errors

import "net/http"

// ErrorCodeConflict means the resource conflicts with an existing one
const ErrorCodeConflict = "conflict"

func NewAppConflictError(field string) AppError {
	return AppError{
		StatusCode: http.StatusConflict,
		Code:       ErrorCodeConflict,
		Field:      field,
	}
}
