package zincsearch

import "net/http"

const (
	ApiDocCreateWithId = "/api/%s/_doc/%s"
	ApiDocDeleteWithId = "/api/%s/_doc/%s"
	ApiSearchEs        = "/es/%s/_search"
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
