package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Title      string             `json:"title" bson:"title" validate:"required"`
	Url        string             `json:"url" bson:"url" validate:"required"`
	Timestamp  string             `json:"timestamp" bson:"timestamp" validate:"required"`
	UserAgent  string             `json:"user_agent" bson:"user_agent" validate:"required"`
	Type       string             `json:"type" bson:"type" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Update_at  time.Time          `json:"updated_at"`
}
