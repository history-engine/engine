package model

import "time"

// MeiliDocHandlerResp 文档操作的统一返回，真实结果使用taskUid查询
type MeiliDocHandlerResp struct {
	TaskUid    int       `json:"taskUid"`
	IndexUid   string    `json:"indexUid"`
	Status     string    `json:"status"`
	Type       string    `json:"type"`
	EnqueuedAt time.Time `json:"enqueuedAt"`
}

// MeiliSearchRequest 查询参数
type MeiliSearchRequest struct {
	Query                string   `json:"q,omitempty"`
	Offset               int      `json:"offset,omitempty"`
	Limit                int      `json:"limit,omitempty"`
	HitsPerPage          int      `json:"hitsPerPage,omitempty"`
	Page                 int      `json:"page,omitempty"`
	Filter               string   `json:"filter,omitempty"`
	Facets               []string `json:"facets,omitempty"`
	AttributesToRetrieve []string `json:"attributesToRetrieve,omitempty"`
}

// MeiliSearchResponse 查询返回参数
type MeiliSearchResponse struct {
	Hits             []MeiliDocument `json:"hits"`
	Query            string          `json:"query"`
	ProcessingTimeMs int             `json:"processingTimeMs"`
	HitsPerPage      int             `json:"hitsPerPage"`
	Page             int             `json:"page"`
	TotalPages       int             `json:"totalPages"`
	TotalHits        int             `json:"totalHits"`
}
