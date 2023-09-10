package entity

import (
	"regexp"
	"strings"
	"time"
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
