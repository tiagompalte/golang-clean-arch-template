package entity

import "time"

type User struct {
	ID        uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID      string
	Name      string
	Email     string
}
