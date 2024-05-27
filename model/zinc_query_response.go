package model

import "time"

// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/query_response.go

// SearchResponse for a query
type SearchResponse struct {
	Took         int                            `json:"took"` // Time it took to generate the response
	TimedOut     bool                           `json:"timed_out"`
	Shards       Shards                         `json:"_shards"`
	Hits         Hits                           `json:"hits"`
	Aggregations map[string]AggregationResponse `json:"aggregations,omitempty"`
	Error        string                         `json:"error,omitempty"`
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index     string                 `json:"_index"`
	Type      string                 `json:"_type"`
	ID        string                 `json:"_id"`
	Score     float64                `json:"_score"`
	Timestamp time.Time              `json:"@timestamp"`
	Source    interface{}            `json:"_source,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Highlight map[string]interface{} `json:"highlight,omitempty"`
}

type Total struct {
	Value int `json:"value"` // Count of documents returned
}

type AggregationResponse struct {
	Value    interface{} `json:"value,omitempty"`
	Buckets  interface{} `json:"buckets,omitempty"`  // slice or map
	Interval string      `json:"interval,omitempty"` // support for auto_date_histogram_aggregation
}
