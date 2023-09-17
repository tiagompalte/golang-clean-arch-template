package errors

import "net/http"

// ErrorCodeBadGateway means that an HTTP Bad Gateway occurred
const ErrorCodeBadGateway = "bad-gateway"

func NewAppBadGatewayError() AppError {
	return AppError{
		StatusCode: http.StatusBadGateway,
		Code:       ErrorCodeBadGateway,
	}
}
