package zincsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"net/http"
	"time"
)

// EsSearch Search V2 ES语法兼容
func EsSearch(userId int64, search model.SearchPage) (resp model.ZincSearchResponse, err error) {
	query := model.ZincQuery{
		Size: search.Limit,
		From: (search.Page - 1) * search.Limit,
		Sort: []string{"-_score"},
		Aggregations: map[string]model.Aggregations{
			"histogram": {
				DateHistogram: &model.AggregationDateHistogram{
					Field:    "@timestamp",
					Interval: "1d",
				},
			},
		},
		Query: model.Query{
			Bool: &model.BoolQuery{
				Must: []*model.Query{
					{
						Range: map[string]*model.RangeQuery{
							"@timestamp": {
								GTE:    search.StartTime.Format(time.RFC3339),
								LT:     search.EndTime.Format(time.RFC3339),
								Format: time.RFC3339,
							},
						},
					},
				},
				Should: []*model.Query{
					{
						Match: map[string]*model.MatchQuery{
							"title": {
								Query: search.Query,
							},
						},
					},
					{
						Match: map[string]*model.MatchQuery{
							"excerpt": {
								Query: search.Query,
							},
						},
					},
					{
						Match: map[string]*model.MatchQuery{
							"content": {
								Query: search.Query,
							},
						},
					},
				},
			},
		},
		Source: []string{"ID", "content", "excerpt"},
	}

	api := fmt.Sprintf(ApiSearchEs, IndexName(userId))
	body, err := SendRequest(api, http.MethodPost, query)
	if err != nil {
		return
	}

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
