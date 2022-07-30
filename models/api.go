package models

type Api struct {
	Base           `json:,inline`
	ApiType        string `json:"api_type" bson:"api_type" validate:"required"`
	Method         string `json:"method" bson:"method"`
	PathUrl        string `json:"path_url" bson:"path_url"`
	RequestBody    string `json:"request_body" bson:"request_body"`
	ResponseBody   string `json:"response_body" bson:"response_body"`
	RequestHeader  string `json:"request_header" bson:"request_header"`
	ResponseHeader string `json:"response_header" bson:"response_header" `
	Success        bool   `json:"success" bson:"success" `
	Status         string `json:"status" bson:"status" `
	Duration       int    `json:"duration" bson:"duration" `
	ErrorMessage   string `json:"error_message" bson:"error_message" `
}
