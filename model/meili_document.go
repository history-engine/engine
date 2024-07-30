package model

// MeiliDocument MeiliSearch 文档，没有固定格式，以实际存储存入的为准
type MeiliDocument struct {
	Id      string `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Title   string `json:"title,omitempty"`
	Excerpt string `json:"excerpt,omitempty"`
	Content string `json:"content,omitempty"`
	Code    string `json:"code"`
}
