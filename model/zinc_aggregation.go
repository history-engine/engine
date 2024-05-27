package model

// https://github.com/zincsearch/zincsearch/blob/main/pkg/bluge/aggregation/histogram.go

type HistogramBound struct {
	Min float64 `json:"min"` // minimum
	Max float64 `json:"max"` // maximum
}
