package zincsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zincsearch/zincsearch/pkg/meta"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"io"
	"net/http"
	"time"
)

// EsSearch Search V2 ES语法兼容
func EsSearch(search model.SearchPage) ([]string, error) {
	query := meta.ZincQueryForSDK{
		Size: search.Limit,
		From: 0,
		Sort: []string{"-@timestamp"},
		Aggregations: map[string]meta.Aggregations{
			"histogram": {
				DateHistogram: &meta.AggregationDateHistogram{
					Field:    "@timestamp",
					Interval: "1d",
				},
			},
		},
		Query: meta.QueryForSDK{
			Bool: &meta.BoolQueryForSDK{
				Must: []*meta.QueryForSDK{
					{
						Range: map[string]*meta.RangeQueryForSDK{
							"@timestamp": {
								GTE:    search.StartTime.Format(time.RFC3339),
								LT:     search.EndTime.Format(time.RFC3339),
								Format: time.RFC3339,
							},
						},
					},
					{
						QueryString: &meta.QueryStringQuery{
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
	resp := &meta.SearchResponse{}
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
