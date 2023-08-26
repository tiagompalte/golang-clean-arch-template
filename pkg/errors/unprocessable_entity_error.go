package errors

// ErrorCodeUnprocessableEntity means the action could not be processed properly due to invalid data provided
const ErrorCodeUnprocessableEntity = "unprocessable-entity"

func NewAppUnprocessableEntityError(message string) AppError {
	return AppError{
		Code:    ErrorCodeUnprocessableEntity,
		Message: message,
	}
}
