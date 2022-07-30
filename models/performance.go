package models

type TimePerformance struct {
	Base                 `json:,inline`
	PerfType             string `json:"perf_type" bson:"perf_type" validate:"required"`
	DnsTime              int    `json:"dns_time" bson:"dns_time"`
	ConnectTime          int    `json:"connect_time" bson:"connect_time"`
	TtfbTime             int    `json:"ttfb_time" bson:"ttfb_time"`
	ResponseTime         int    `json:"response_time" bson:"response_time"`
	ParseDOMTime         int    `json:"parse_dom_time" bson:"parse_dom_time"`
	DomContentLoadedTime int    `json:"dom_content_loaded_time" bson:"dom_content_loaded_time"`
	TimeToInteractive    int    `json:"time_to_interactive" bson:"time_to_interactive"`
	LoadTime             int    `json:"load_time" bson:"load_time" `
}

type PaintPerformance struct {
	Base                   `json:,inline`
	PerfType               string `json:"perf_type" bson:"perf_type" validate:"required"`
	FirstPaint             int    `json:"first_paint" bson:"first_paint"`
	FirstContentPaint      int    `json:"first_content_paint" bson:"first_contentful_paint"`
	FirstMeaningfulPaint   int    `json:"first_meaningful_paint" bson:"first_meaningful_paint"`
	LargestContentfulPaint int    `json:"largest_contentful_paint" bson:"largest_contentful_paint"`
	FirstInputDelay        int    `json:"first_input_delay" bson:"first_input_delay"`
}
