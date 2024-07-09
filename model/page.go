package model

import "time"

type Page struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	UniqueId  string    `json:"unique_id" db:"unique_id"`
	Version   int       `json:"version" db:"version"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url" db:"url"`
	FullPath  string    `json:"full_path" db:"full_path"`
	FullSize  int       `json:"full_size" db:"full_size"`
	LitePath  string    `json:"lite_path" db:"lite_path"`
	LiteSize  int       `json:"lite_size" db:"lite_size"`
	IndexedAt time.Time `json:"indexed_at" db:"indexed_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type SearchPage struct {
	Query     string    `json:"query" query:"query"`
	Limit     int       `json:"limit" query:"limit"`
	StartTime time.Time `json:"start_time" query:"start_time"`
	EndTime   time.Time `json:"end_time" query:"end_time"`
}

type PageSearchResponse struct {
	Avatar  string `json:"avatar"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Size    int    `json:"size"`
	Url     string `json:"url"`
	Preview string `json:"preview"`
}
