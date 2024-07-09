package zincsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"io"
	"net/http"
	"time"
)

// EsSearch Search V2 ES语法兼容
// todo 返回条数，用于分页
func EsSearch(search model.SearchPage) (map[string]model.ZincDocument, []string, error) {
	query := model.ZincQueryForSDK{
		Size: search.Limit,
		From: 0,
		Sort: []string{"-@timestamp"},
		Aggregations: map[string]model.Aggregations{
			"histogram": {
				DateHistogram: &model.AggregationDateHistogram{
					Field:    "@timestamp",
					Interval: "1d",
				},
			},
		},
		Query: model.QueryForSDK{
			Bool: &model.BoolQueryForSDK{
				Must: []*model.QueryForSDK{
					{
						Range: map[string]*model.RangeQueryForSDK{
							"@timestamp": {
								GTE:    search.StartTime.Format(time.RFC3339),
								LT:     search.EndTime.Format(time.RFC3339),
								Format: time.RFC3339,
							},
						},
					},
					{
						QueryString: &model.QueryStringQuery{
							Query: search.Query,
						},
					},
				},
			},
		},
		Source: []string{"ID", "title", "content", "url", "size"},
	}
	body, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	api := fmt.Sprintf(ApiSearchV2, setting.ZincSearch.Index)
	req, err := http.NewRequest(http.MethodPost, setting.ZincSearch.Host+api, bytes.NewReader(body))
	if err != nil {
		logger.Zap().Error("new request error", zap.Error(err))
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	resp := &model.ZincSearchResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		panic(err)
		logger.Zap().Error("search zinc err", zap.ByteString("resp", body))
		return nil, nil, err
	}

	if resp.Error != "" {
		return nil, nil, errors.New(resp.Error)
	}

	docs := make(map[string]model.ZincDocument, 0)
	ids := make([]string, resp.Hits.Total.Value)
	for k, v := range resp.Hits.Hits {
		ids[k] = v.ID
		if source, ok := v.Source.(map[string]interface{}); ok {
			docs[v.ID] = model.ZincDocument{
				Id:      v.ID,
				Title:   source["title"].(string),
				Content: source["content"].(string),
				Url:     source["url"].(string),
				Size:    int(source["size"].(float64)),
			}
		}

	}

	return docs, ids, nil
}
