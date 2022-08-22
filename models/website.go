package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Website struct {
	// Base           `json:,inline`
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string             `json:"title" bson:"title" validate:"required"`
	Url     string             `json:"url" bson:"url" validate:"required"`
	Creator string             `json:"creator" bson:"creator" validate:"required"`
	Desc    string             `json:"desc" bson:"desc" validate:"required"`
}
