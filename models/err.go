package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JsError struct {
	// Base      `json:,inline bson:,inline`
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

	// jsError
	ErrorType string `json:"errorType" bson:"error_type" validate:"required"`
	Message   string `json:"message" bson:"message" validate:"required"`
	Stack     string `json:"stack" bson:"stack"`
}

type ResourceError struct {
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

	// resourceError
	ErrorType    string  `json:"errorType" bson:"error_type" validate:"required"` //错误类型
	FileName     string  `json:"filename" bson:"file_name"`                       //访问的文件名
	ErrorMessage string  `json:"errorMessage" bson:"error_message"`               //错误信息
	TagName      string  `json:"tagName" bson:"tag_name"`                         //标签名
	Size         string  `json:"size" bson:"size"`                                // 资源大小
	Time         float64 `json:"time" bson:"time"`                                // 请求时间
}

// unused
type BlankError struct {
	// Base        `json:,inline bson:,inline`
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

	// blankError
	ErrorType   string `json:"errorType" bson:"error_type" validate:"required"` //错误类型
	EmptyPoints string `json:"emptyPoints" bson:"empty_points"`                 //空点
	Screen      string `json:"screen" bson:"screen"`                            //分辨率
	ViewPoint   string `json:"viewPoint" bson:"view_point"`                     //视口              /
}
