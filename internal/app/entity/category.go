package entity

import (
	"regexp"
	"strings"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type Category struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	slug      string
	Name      string
	UserID    uint32
}

var regexSlugReplace = regexp.MustCompile(`\s+`)

func (c *Category) GetSlug() string {
	if c.slug == "" {
		c.slug = regexSlugReplace.ReplaceAllString(strings.ToLower(c.Name), "-")
	}
	return c.slug
}

func (c *Category) SetSlug(slug string) {
	c.slug = slug
}

func (c Category) ValidateNew() error {
	aggrErr := errors.NewAggregatedError()

	if c.Name == "" {
		aggrErr.Add(errors.NewEmptyParameterError("name"))
	}

	if c.UserID == 0 {
		aggrErr.Add(errors.NewEmptyParameterError("user_id"))
	}

	if aggrErr.Len() > 0 {
		return errors.Wrap(aggrErr)
	}

	return nil
}
