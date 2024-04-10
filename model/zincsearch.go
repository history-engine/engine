package model

import "time"

type ZincDocument struct {
	FilePath string `json:"file_path"` // todo 不需要
	Url      string `json:"url"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Size     int    `json:"size"`
}

type ZincSearch struct {
	SearchType string          `json:"search_type"`
	Query      ZincSearchQuery `json:"query"`
	SortFields []string        `json:"sort_fields"`
	From       int             `json:"from"`
	MaxResults int             `json:"max_results"`
	Source     []string        `json:"_source"`
}

type ZincSearchQuery struct {
	Term      string    `json:"term"`
	Field     string    `json:"field"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ZincResponse struct {
	Took     int      `json:"took"`
	TimedOut bool     `json:"timed_out"`
	MaxScore int      `json:"max_score"`
	Hits     ZincHits `json:"hits"`
}

type ZincHits struct {
}
