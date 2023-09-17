package errors

import "net/http"

// ErrorCodeUnauthorized means the user wasn't identified
const ErrorCodeUnauthorized = "unauthorized"

func NewAppUnauthorizedError() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeUnauthorized,
	}
}
