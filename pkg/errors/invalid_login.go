package errors

import (
	"net/http"
)

// ErrorCodeInvalidLogin means that path is empty
const ErrorCodeInvalidLogin = "invalid-login"

func NewInvalidLoginError() AppError {
	return AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidLogin,
	}
}
