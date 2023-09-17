package errors

import (
	"net/http"

	pkg "github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

// ErrorCodeEmptyPath means that path is empty
const ErrorCodeInvalidLogin = "invalid-login"

func NewInvalidLoginError() pkg.AppError {
	return pkg.AppError{
		StatusCode: http.StatusUnauthorized,
		Code:       ErrorCodeInvalidLogin,
	}
}
