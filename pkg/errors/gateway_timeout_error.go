package errors

import "net/http"

// ErrorCodeBadGateway means that server was acting as a gateway or proxy and did not receive a timely response from the upstream server
const ErrorCodeGatewayTimeout = "gateway-timeout"

func NewAppGatewayTimeoutError() AppError {
	return AppError{
		StatusCode: http.StatusGatewayTimeout,
		Code:       ErrorCodeGatewayTimeout,
	}
}
