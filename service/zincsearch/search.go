package zincsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"io"
	"net/http"
	"time"
)

// EsSearch Search V2 ES语法兼容
func EsSearch(search model.SearchPage) ([]string, error) {
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
		Source: []string{"ID"},
	}
	body, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	api := fmt.Sprintf(ApiSearchV2, setting.ZincSearch.Index)
	req, err := http.NewRequest(http.MethodPost, setting.ZincSearch.Host+api, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	resp := &model.SearchResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		panic(err)
	}

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	ids := make([]string, resp.Hits.Total.Value)
	for k, v := range resp.Hits.Hits {
		ids[k] = v.ID
	}

	return ids, nil
}
