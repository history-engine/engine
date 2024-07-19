package readability

type Article struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Excerpt     string `json:"excerpt"`
	Byline      string `json:"byline"`
	Dir         string `json:"dir"`
	Lang        string `json:"lang"`
	Content     string `json:"content"`
	TextContent string `json:"textContent"`
	Length      int    `json:"length"`
	SiteName    string `json:"siteName"`
}

type Readability interface {
	Parse(path string) (*Article, error)
	ParseContent(content []byte) (*Article, error)
	ExtractSingleFileUrl(content []byte) string
}
