package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Log struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at"`
	Level     string        `bson:"level"`
	Message   any           `bson:"message"`
}
