package errors

import (
	pkg "github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

// ErrorCodeEmptyParameter means that parameter is empty
const ErrorCodeEmptyParameter = "empty-parameter"

func NewEmptyParameterError(field string) pkg.AppError {
	return pkg.AppError{
		Code:  ErrorCodeEmptyParameter,
		Field: field,
	}
}
