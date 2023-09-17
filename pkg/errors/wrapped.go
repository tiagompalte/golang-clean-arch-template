package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type errorWrapper interface {
	error
	Cause() error
}

func Wrap(err error, messages ...string) error {
	if err != nil {
		// Gets caller function path.
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		funcRef := runtime.FuncForPC(pc[0])
		pathArr := strings.Split(funcRef.Name(), "/")
		path := pathArr[len(pathArr)-1]

		return WrappedError{
			originalError: err,
			path:          path,
			messages:      messages,
		}
	}

	return nil
}

func Cause(err error) error {
	wrappedErr, ok := err.(errorWrapper)
	if !ok {
		return err
	}
	return wrappedErr.Cause()
}

type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

func (e WrappedError) Error() string {
	errPath := e.path
	if len(e.messages) > 0 {

		errPath += ": "
		for _, message := range e.messages {
			errPath += message + "; "
		}
	}

	return fmt.Sprintf("%s → %v", errPath, e.originalError)
}

func (e WrappedError) Cause() error {
	if e.originalError != nil {
		originalError, ok := (e.originalError).(errorWrapper)
		if ok {
			return originalError.Cause()
		}
	}

	return e.originalError
}
