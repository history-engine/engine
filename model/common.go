package model

type ApiResponse struct {
	Code    MsgCode `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"` // 任意时候都要返回对象, 防止新增字段
}

// PageResponse 分页返回
type PageResponse[T any] struct {
	Total int `json:"total,omitempty"`
	Data  T   `json:"data,omitempty"`
}
