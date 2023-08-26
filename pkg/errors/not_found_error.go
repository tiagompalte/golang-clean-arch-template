package errors

// ErrorCodeNotFound means the request resource was not found
const ErrorCodeNotFound = "not-found"

func NewAppNotFoundError(field string) AppError {
	return AppError{
		Code:  ErrorCodeNotFound,
		Field: field,
	}
}
