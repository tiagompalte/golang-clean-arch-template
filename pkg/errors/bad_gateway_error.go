package errors

// ErrorCodeBadGateway means that an HTTP Bad Gateway occurred
const ErrorCodeBadGateway = "bad-gateway"

func NewAppBadGatewayError() AppError {
	return AppError{
		Code: ErrorCodeBadGateway,
	}
}
