package model

type ZincDocument struct {
	FilePath string `json:"file_path"` // todo 不需要
	Url      string `json:"url"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Size     int    `json:"size"`
}
