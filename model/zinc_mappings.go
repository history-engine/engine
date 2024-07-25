package model

import "sync"

// https://github.com/zincsearch/zincsearch/blob/main/pkg/meta/mappings.go

type Mappings struct {
	Properties map[string]Property `json:"properties,omitempty"`
	lock       sync.RWMutex
}

type Property struct {
	Type           string `json:"type"` // text, keyword, date, numeric, boolean, geo_point
	Analyzer       string `json:"analyzer,omitempty"`
	SearchAnalyzer string `json:"search_analyzer,omitempty"`
	Format         string `json:"format,omitempty"`    // date format yyyy-MM-dd HH:mm:ss || yyyy-MM-dd || epoch_millis
	TimeZone       string `json:"time_zone,omitempty"` // date format time_zone
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Aggregatable   bool   `json:"aggregatable"`
	Highlightable  bool   `json:"highlightable"`
	// Fields allow the same string value to be indexed in multiple ways for different purposes,
	// such as one field for search and a multi-field for sorting and aggregations,
	// or the same string value analyzed by different analyzers.
	// If the Fields property is defined within a sub-field, it will be ignored.
	//
	// Currently, only "text" fields support the Fields parameter.
	Fields map[string]Property `json:"fields,omitempty"`
}
