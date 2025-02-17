package errors

import "net/http"

// ErrorCodeConcurrencyRepository indicates that the version informed is not the same from database
const ErrorCodeConcurrencyRepository = "concurrency-repository"

func NewAppConcurrencyRepositoryError() AppError {
	return AppError{
		StatusCode: http.StatusBadRequest,
		Code:       ErrorCodeConcurrencyRepository,
	}
}
