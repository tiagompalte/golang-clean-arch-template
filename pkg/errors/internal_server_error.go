package errors

import "net/http"

// ErrorCodeInternalServerError means an unexpected condition was encountered and no more specific message is suitable
const ErrorCodeInternalServerError = "internal-server-error"

func NewAppInternalServerError() AppError {
	return AppError{
		StatusCode: http.StatusInternalServerError,
		Code:       ErrorCodeInternalServerError,
	}
}
