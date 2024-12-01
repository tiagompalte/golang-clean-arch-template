package errors

import (
	"net/http"
)

// ErrorCodeInvalidUser means that user is invalid
const ErrorCodeInvalidUser = "invalid-user"

func NewInvalidUserError() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidUser,
	}
}
