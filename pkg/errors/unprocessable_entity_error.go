package errors

import "net/http"

// ErrorCodeUnprocessableEntity means the action could not be processed properly due to invalid data provided
const ErrorCodeUnprocessableEntity = "unprocessable-entity"

func NewAppUnprocessableEntityError(message string) AppError {
	return AppError{
		StatusCode: http.StatusUnprocessableEntity,
		Code:       ErrorCodeUnprocessableEntity,
		Message:    message,
	}
}
