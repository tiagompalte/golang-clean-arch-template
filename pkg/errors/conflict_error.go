package errors

// ErrorCodeConflict means the resource conflicts with an existing one
const ErrorCodeConflict = "conflict"

func NewAppConflictError(field string) AppError {
	return AppError{
		Code:  ErrorCodeConflict,
		Field: field,
	}
}
