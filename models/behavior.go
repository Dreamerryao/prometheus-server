package models

type PvBehavior struct {
	Base         `json:,inline`
	BehaviorType string `json:"behavior_type" bson:"behavior_type" validate:"required"`
	PageUrl      string `json:"page_url" bson:"page_url" validate:"required"`
	Uuid         string `json:"uuid" bson:"uuid" validate:"required"`
}

type StayBehavior struct {
	Base         `json:,inline`
	BehaviorType string `json:"behavior_type" bson:"behavior_type" validate:"required"`
	PageUrl      string `json:"page_url" bson:"page_url" validate:"required"`
	Uuid         string `json:"uuid" bson:"uuid" validate:"required"`
	StayTime     int    `json:"stay_time" bson:"stay_time" validate:"required"`
}
