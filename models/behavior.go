package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PvBehavior struct {
	// Base         `json:,inline bson:,inline`
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Url       string             `json:"url" bson:"url" validate:"required"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
	UserAgent string             `json:"userAgent" bson:"user_agent"`
	Referrer  string             `json:"referrer" bson:"referrer"`
	// NavigationType int                `json:"navigationType" bson:"navigation_type"`
	Type       string    `json:"type" bson:"type" validate:"required"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Update_at  time.Time `json:"updated_at" bson:"updated_at"`

	//pvBehavior
	BehaviorType string `json:"behavior_type" bson:"behavior_type" validate:"required"`
	PageUrl      string `json:"page_url" bson:"page_url" validate:"required"`
	Uuid         string `json:"uuid" bson:"uuid" validate:"required"`
}

type StayBehavior struct {
	// Base         `json:,inline`
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Url       string             `json:"url" bson:"url" validate:"required"`
	Timestamp string             `json:"timestamp" bson:"timestamp"`
	UserAgent string             `json:"userAgent" bson:"user_agent"`
	Referrer  string             `json:"referrer" bson:"referrer"`
	// NavigationType int                `json:"navigationType" bson:"navigation_type"`
	Type       string    `json:"type" bson:"type" validate:"required"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Update_at  time.Time `json:"updated_at" bson:"updated_at"`

	// stayBehavior
	BehaviorType string  `json:"behavior_type" bson:"behavior_type" validate:"required"`
	PageUrl      string  `json:"page_url" bson:"page_url" validate:"required"`
	Uuid         string  `json:"uuid" bson:"uuid" validate:"required"`
	StayTime     float64 `json:"stay_time" bson:"stay_time" validate:"required"`
}
