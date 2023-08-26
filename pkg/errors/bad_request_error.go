package errors

// ErrorCodeBadRequest indicates that the request sent to the server is invalid or corrupted
const ErrorCodeBadRequest = "bad-request"

func NewAppBadRequestError(message string) AppError {
	return AppError{
		Code:    ErrorCodeBadRequest,
		Message: message,
	}
}
