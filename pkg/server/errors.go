package server

import (
	"fmt"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

func HandleError(handle Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handle(w, r)
		if err != nil {
			prepareError(w, r, err)
		}
	}
}

type ErrorResponse struct {
	Code    string `json:"code,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorResponseWrapper struct {
	Errors []ErrorResponse `json:"errors,omitempty"`
}

func NewErrorResponseWrapper(errs ...error) ErrorResponseWrapper {
	listErrors := make([]ErrorResponse, len(errs))
	for i := range errs {
		switch e := errs[i].(type) {
		case errors.AppError:
			listErrors[i].Code = e.Code
			listErrors[i].Field = e.Field
			listErrors[i].Message = e.Message
		default:
			listErrors[i].Code = errors.ErrorCodeInternalServerError
		}
	}
	return ErrorResponseWrapper{
		Errors: listErrors,
	}
}

func prepareError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err.Error())

	var respError error
	switch e := err.(type) {

	case errors.WrappedError:
		prepareError(w, r, e.Cause())

	case errors.AppError:
		respError = RespondError(w, e)

	case errors.AggregatedError:
		respError = RespondAggregateError(w, http.StatusUnprocessableEntity, e)

	default:
		respError = RespondError(w, errors.NewAppInternalServerError())
	}

	if respError != nil {
		prepareError(w, r, respError)
	}
}
