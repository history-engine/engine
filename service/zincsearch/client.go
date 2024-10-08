package zincsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"history-engine/engine/setting"
	"io"
	"net/http"
)

// todo ZincSearch 调用封装，可以考虑封装为独立的sdk，search包直接调用sdk

const (
	ApiDocCreateWithId       = "/api/%s/_doc/%s"
	ApiDocDeleteWithId       = "/api/%s/_doc/%s"
	ApiDocGetWithId          = "/api/%s/_doc/%s"
	ApiIndexCreateUpdateList = "/api/index"
	ApiSearchEs              = "/es/%s/_search"
)

var client *http.Client

func init() {
	client = http.DefaultClient
}

func SetClient(c *http.Client) {
	client = c
}

func SendRequest(ctx context.Context, api, method string, data any) ([]byte, error) {
	var err error
	var content []byte
	if data != nil {
		content, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, setting.ZincSearch.Host+api, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
