package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TimePerformance struct {
	// Base                 `json:,inline`
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

	// timePerformance
	PerfType             string  `json:"perfType" bson:"perf_type" validate:"required"`
	DnsTime              float64 `json:"dnsTime" bson:"dns_time"`
	ConnectTime          float64 `json:"connectTime" bson:"connect_time"`
	TtfbTime             float64 `json:"ttfbTime" bson:"ttfb_time"`
	ResponseTime         float64 `json:"responseTime" bson:"response_time"`
	ParseDOMTime         float64 `json:"parseDomTime" bson:"parse_dom_time"`
	DomContentLoadedTime float64 `json:"domContentLoadedTime" bson:"dom_content_loaded_time"`
	TimeToInteractive    float64 `json:"timeToInteractive" bson:"time_to_interactive"`
	LoadTime             float64 `json:"loadTime" bson:"load_time" `
}

type PaintPerformance struct {
	// Base                   `json:,inline`
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

	// paintPerformance
	PerfType               string  `json:"perfType" bson:"perf_type" validate:"required"`
	FirstPaint             float64 `json:"firstPaint" bson:"first_paint"`
	FirstContentPaint      float64 `json:"firstContentPaint" bson:"first_contentful_paint"`
	FirstMeaningfulPaint   float64 `json:"firstMeaningfulPaint" bson:"first_meaningful_paint"`
	LargestContentfulPaint float64 `json:"largestContentfulPaint" bson:"largest_contentful_paint"`
	FirstInputDelay        float64 `json:"firstInputDelay" bson:"first_input_delay"`
}
