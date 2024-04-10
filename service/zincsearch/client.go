package zincsearch

import "net/http"

const (
	ApiIndex    = "/api/%s/_doc/%s"
	ApiSearchV2 = "/es/%s/_search"
)

var client *http.Client = http.DefaultClient

func init() {
	// todo 深入自定义http client
}

func SetClient(c *http.Client) {
	client = c
}

func SendRequest(api, method string, body []byte) {
	// todo
}
