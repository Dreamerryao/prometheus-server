package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Api struct {
	// Base           `json:,inline`
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

	// api
	ApiType        string  `json:"apiType" bson:"api_type" validate:"required"`
	Method         string  `json:"method" bson:"method"`
	PathUrl        string  `json:"pathUrl" bson:"path_url"`
	RequestBody    string  `json:"requestBody" bson:"request_body"`
	ResponseBody   string  `json:"responseBody" bson:"response_body"`
	RequestHeader  string  `json:"requestHeader" bson:"request_header"`
	ResponseHeader string  `json:"responseHeader" bson:"response_header" `
	Success        bool    `json:"success" bson:"success" `
	Status         float64 `json:"status" bson:"status" `
	Duration       float64 `json:"duration" bson:"duration" `
	ErrorMessage   string  `json:"errorMessage" bson:"error_message" `
}
