package model

import "time"

type SearchPage struct {
	Query     string    `json:"query" query:"query"`
	Page      int       `json:"page" query:"page"`
	Limit     int       `json:"limit" query:"limit"`
	StartTime time.Time `json:"start_time" query:"start_time"`
	EndTime   time.Time `json:"end_time" query:"end_time"`
}

type PageSearch struct {
	Avatar  string `json:"avatar"`
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
	Size    int    `json:"size"`
	Url     string `json:"url"`
	Preview string `json:"preview"`
	Version int    `json:"version"`
}

type SearchResponse struct {
	Total int `json:"total"`
	Pages any `json:"pages"`
}
