package model

import "time"

type Page struct {
	Id        int64     `json:"id" db:"id"`
	UserId    int64     `json:"user_id" db:"user_id"`
	UniqueId  string    `json:"unique_id" db:"unique_id"`
	Version   int       `json:"version" db:"version"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url" db:"url"`
	Path      string    `json:"path" db:"path"`
	Size      int       `json:"size" db:"size"`
	IndexedAt time.Time `json:"indexed_at" db:"indexed_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

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
