package meilisearch

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"net/http"
)

func Search(ctx context.Context, userId int64, search model.SearchRequest) (resp model.MeiliSearchResponse, err error) {
	query := model.MeiliSearchRequest{
		Query:       search.Query,
		HitsPerPage: search.Limit,
		Page:        search.Page,
	}
	api := fmt.Sprintf(ApiSearch, IndexName(userId))
	_, body, err := SendRequest(ctx, api, http.MethodPost, query)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.Zap().Error("search zinc err", zap.ByteString("resp", body))
		return
	}

	return resp, nil
}
