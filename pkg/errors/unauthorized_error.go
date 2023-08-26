package errors

// ErrorCodeUnauthorized means the user wasn't identified
const ErrorCodeUnauthorized = "unauthorized"

func NewAppUnauthorizedError() AppError {
	return AppError{
		Code: ErrorCodeUnauthorized,
	}
}
