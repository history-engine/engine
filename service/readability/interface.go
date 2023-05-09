package readability

type Article struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Byline      string `json:"byline"`
	Dir         string `json:"dir"`
	Lang        string `json:"lang"`
	Content     string `json:"content"`
	TextContent string `json:"textContent"`
	Length      int    `json:"length"`
	SiteName    string `json:"siteName"`
}

type Readability interface {
	Parse(path string) *Article
	ParseContent(content []byte) *Article
	ExtractSingleFileUrl(content []byte) string
}
