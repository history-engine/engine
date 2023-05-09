package model

type ApiResponse struct {
	Code    MsgCode `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"` // 任意时候都要返回对象, 防止新增字段
}
