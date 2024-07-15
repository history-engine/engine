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
func EsSearch(search model.SearchPage) (resp model.ZincSearchResponse, err error) {
	query := model.ZincQueryForSDK{
		Size: search.Limit,
		From: (search.Page - 1) * search.Limit,
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

	api := fmt.Sprintf(ApiSearchEs, setting.ZincSearch.Index)
	req, err := http.NewRequest(http.MethodPost, setting.ZincSearch.Host+api, bytes.NewReader(body))
	if err != nil {
		logger.Zap().Error("new request error", zap.Error(err))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.Zap().Error("search zinc err", zap.ByteString("resp", body))
		return
	}

	if resp.Error != "" {
		return resp, errors.New(resp.Error)
	}

	return resp, nil
}
