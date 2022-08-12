package models

type TimePerformance struct {
	Base                 `json:,inline`
	PerfType             string  `json:"perf_type" bson:"perf_type" validate:"required"`
	DnsTime              float64 `json:"dns_time" bson:"dns_time"`
	ConnectTime          float64 `json:"connect_time" bson:"connect_time"`
	TtfbTime             float64 `json:"ttfb_time" bson:"ttfb_time"`
	ResponseTime         float64 `json:"response_time" bson:"response_time"`
	ParseDOMTime         float64 `json:"parse_dom_time" bson:"parse_dom_time"`
	DomContentLoadedTime float64 `json:"dom_content_loaded_time" bson:"dom_content_loaded_time"`
	TimeToInteractive    float64 `json:"time_to_interactive" bson:"time_to_interactive"`
	LoadTime             float64 `json:"load_time" bson:"load_time" `
}

type PaintPerformance struct {
	Base                   `json:,inline`
	PerfType               string  `json:"perf_type" bson:"perf_type" validate:"required"`
	FirstPaint             float64 `json:"first_paint" bson:"first_paint"`
	FirstContentPaint      float64 `json:"first_content_paint" bson:"first_contentful_paint"`
	FirstMeaningfulPaint   float64 `json:"first_meaningful_paint" bson:"first_meaningful_paint"`
	LargestContentfulPaint float64 `json:"largest_contentful_paint" bson:"largest_contentful_paint"`
	FirstInputDelay        float64 `json:"first_input_delay" bson:"first_input_delay"`
}
