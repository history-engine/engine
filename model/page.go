package model

import (
	"io"
	"time"
)

// SearchRequest WebUI搜索请求参数
type SearchRequest struct {
	Query     string    `json:"query" query:"query"`
	Page      int       `json:"page" query:"page"`
	Limit     int       `json:"limit" query:"limit"`
	StartTime time.Time `json:"start_time" query:"start_time"`
	EndTime   time.Time `json:"end_time" query:"end_time"`
}

// SearchResponse WebUI搜索结果参数
type SearchResponse struct {
	Total int                `json:"total"` // 总条数
	Pages []SearchResultPage `json:"pages"` // 页面数据
}

// SearchResultPage 返回给WebUI的页面参数
type SearchResultPage struct {
	DocId     string    `json:"doc_id"`
	UniqueId  string    `json:"unique_id"`
	Version   int       `json:"version"`
	Avatar    string    `json:"avatar"`
	Title     string    `json:"title"`
	Excerpt   string    `json:"excerpt"`
	Content   string    `json:"content"`
	Size      int       `json:"size"`
	Url       string    `json:"url"`
	Preview   string    `json:"preview"`
	CreatedAt time.Time `json:"created_at"`
}

type PageParse struct {
	Id       int64
	UserId   int64
	UniqueId string
	Version  int
}

type HtmlInfo struct {
	Host     string
	Url      string
	Suffix   string
	Sha1     string
	Size     int
	UserId   int64
	IoReader io.ReadCloser
}
