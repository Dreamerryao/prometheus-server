package models

type JsError struct {
	Base      `json:,inline`
	ErrorType string `json:"error_type" bson:"error_type" validate:"required"`
	Message   string `json:"message" bson:"message" validate:"required"`
	Stack     string `json:"stack" bson:"stack"`
}

type ResourceError struct {
	Base         `json:,inline`
	ErrorType    string `json:"error_type" bson:"error_type" validate:"required"` //错误类型
	FileName     string `json:"file_name" bson:"file_name"`                       //访问的文件名
	ErrorMessage string `json:"error_message" bson:"error_message"`               //错误信息
	TagName      string `json:"tag_name" bson:"tag_name"`                         //标签名
	Size         string `json:"size" bson:"size"`                                 // 资源大小
	Time         string `json:"time" bson:"time"`                                 // 请求时间
}

// unused
type BlankError struct {
	Base        `json:,inline`
	ErrorType   string `json:"error_type" bson:"error_type" validate:"required"` //错误类型
	EmptyPoints string `json:"empty_points" bson:"empty_points"`                 //空点
	Screen      string `json:"screen" bson:"screen"`                             //分辨率
	ViewPoint   string `json:"view_point" bson:"view_point"`                     //视口              /
}
