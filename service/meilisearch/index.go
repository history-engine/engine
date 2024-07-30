package meilisearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"net/http"
)

func CreateIndex(ctx context.Context, userId int64) error {
	mi := model.MeiliIndex{
		Uid:        IndexName(userId),
		PrimaryKey: "id",
	}
	content, err := SendRequest(ctx, ApiIndexCreate, http.MethodPost, mi)
	if err != nil {
		return err
	}

	mdr := &model.MeiliDocHandlerResp{}
	err = json.Unmarshal(content, mdr)
	if err == nil && mdr.Status != "enqueued" {
		return errors.New(mdr.Status)
	}

	return err
}

func IndexName(userId int64) string {
	return fmt.Sprintf("%s_%s_%d", setting.Search.Prefix, setting.Common.Env, userId)
}
