package errors

import (
	"fmt"
)

type AppError struct {
	StatusCode int
	Code       string
	Field      string
	Message    string
}

func (e AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	if e.Field == "" {
		return fmt.Sprintf("error code %s", e.Code)
	}

	return fmt.Sprintf("error code %s on field %s", e.Code, e.Field)
}

func IsAppError(err error, code string) bool {
	cause := Cause(err)
	if appErr, ok := cause.(AppError); ok {
		return appErr.Code == code
	}
	return false
}
