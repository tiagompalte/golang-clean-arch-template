package entity

import "time"

type Log struct {
	ID        any       `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	Level     string    `bson:"level"`
	Message   string    `bson:"message"`
}
