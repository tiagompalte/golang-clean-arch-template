package entity

import (
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type User struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   uint32
	UUID      string
	Name      string
	Email     string
}

func (u User) ValidateNew() error {
	aggrErr := errors.NewAggregatedError()

	if u.Name == "" {
		aggrErr.Add(errors.NewEmptyParameterError("name"))
	}
	if u.Email == "" {
		aggrErr.Add(errors.NewEmptyParameterError("email"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}
