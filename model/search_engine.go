package model

type SearchEngineDocument struct {
	Id      string `json:"id,omitempty"`
	Url     string `json:"url,omitempty"`
	Title   string `json:"title,omitempty"`
	Excerpt string `json:"excerpt,omitempty"`
	Content string `json:"content,omitempty"`
}
