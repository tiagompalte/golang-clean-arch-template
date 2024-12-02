package entity

import (
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type Task struct {
	ID          uint32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UUID        string
	Name        string
	Description string
	Done        bool
	Category    Category
	UserID      uint32
}

func (t Task) ValidateNew() error {
	aggrErr := errors.NewAggregatedError()

	if t.Name == "" {
		aggrErr.Add(errors.NewEmptyParameterError("name"))
	}
	if t.Description == "" {
		aggrErr.Add(errors.NewEmptyParameterError("description"))
	}
	if t.UserID == 0 {
		aggrErr.Add(errors.NewEmptyParameterError("user_id"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}
