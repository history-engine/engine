package meilisearch

import (
	"bytes"
	"context"
	"encoding/json"
	"history-engine/engine/setting"
	"io"
	"net/http"
)

// todo MeiliSearch 调用封装，理论上整个包都可以删除掉，search包改为调用官方sdk

const (
	ApiDocCreate    = "/indexes/%s/documents"
	ApiDocDelete    = "/indexes/%s/documents/%s"
	ApiDocGetWithId = "/indexes/%s/documents/%s"
	ApiIndexCreate  = "/indexes"
	ApiSearch       = "/indexes/%s/search"
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

	req, err := http.NewRequestWithContext(ctx, method, setting.MeiliSearch.Host+api, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+setting.MeiliSearch.MasterKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
