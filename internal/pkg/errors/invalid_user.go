package errors

import (
	"net/http"

	pkg "github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

// ErrorCodeInvalidUser means that user is invalid
const ErrorCodeInvalidUser = "invalid-user"

func NewInvalidUserError() pkg.AppError {
	return pkg.AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidUser,
	}
}
