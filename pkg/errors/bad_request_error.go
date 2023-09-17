package errors

import "net/http"

// ErrorCodeBadRequest indicates that the request sent to the server is invalid or corrupted
const ErrorCodeBadRequest = "bad-request"

func NewAppBadRequestError(message string) AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeBadRequest,
		Message:    message,
	}
}
