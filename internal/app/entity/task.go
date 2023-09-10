package entity

import (
	"time"
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
