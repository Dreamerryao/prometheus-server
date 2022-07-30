package models

type Base struct {
	Title     string `json:"title" bson:"title"`
	Url       string `json:"url" bson:"url"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
	UserAgent string `json:"user_agent" bson:"user_agent"`
	Type      string `json:"type" bson:"type"`
	Ip        string `json:"ip" bson:"ip"`
}
