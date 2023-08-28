package errors

import pkg "github.com/tiagompalte/golang-clean-arch-template/pkg/errors"

// ErrorCodeEmptyPath means that path is empty
const ErrorCodeEmptyPath = "empty-path"

func NewEmptyPathError(field string) pkg.AppError {
	return pkg.AppError{
		Code:  ErrorCodeEmptyPath,
		Field: field,
	}
}
