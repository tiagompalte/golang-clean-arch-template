package errors

import (
	"fmt"
)

type AggregatedError []error

func NewAggregatedError(errs ...error) AggregatedError {
	return errs
}

func (e AggregatedError) Error() string {
	var message string
	for _, err := range e {
		message += fmt.Sprintf("- %s\n", err.Error())
	}

	return message
}

func (e AggregatedError) Len() int {
	return len(e)
}

func (e *AggregatedError) Add(err error) {
	*e = append(*e, err)
}

func (e *AggregatedError) AddList(err []error) {
	*e = append(*e, err...)
}

func (e AggregatedError) Return() error {
	if e.Len() > 0 {
		return Wrap(e)
	}
	return nil
}
