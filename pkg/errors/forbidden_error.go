package errors

import "net/http"

// ErrorCodeForbidden means user not having the necessary permissions for a resource
const ErrorCodeForbidden = "forbidden"

func NewAppForbiddenError() AppError {
	return AppError{
		StatusCode: http.StatusForbidden,
		Code:       ErrorCodeForbidden,
	}
}
