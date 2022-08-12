package models

type Api struct {
	Base           `json:,inline`
	ApiType        string  `json:"apiType" bson:"api_type" validate:"required"`
	Method         string  `json:"method" bson:"method"`
	PathUrl        string  `json:"pathUrl" bson:"path_url"`
	RequestBody    string  `json:"requestBody" bson:"request_body"`
	ResponseBody   string  `json:"responseBody" bson:"response_body"`
	RequestHeader  string  `json:"requestHeader" bson:"request_header"`
	ResponseHeader string  `json:"responseHeader" bson:"response_header" `
	Success        bool    `json:"success" bson:"success" `
	Status         string  `json:"status" bson:"status" `
	Duration       float64 `json:"duration" bson:"duration" `
	ErrorMessage   string  `json:"error_message" bson:"error_message" `
}
