package meilisearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"net/http"
)

func GetDocument(ctx context.Context, userId int64, docId string) (*model.MeiliDocument, error) {
	api := fmt.Sprintf(ApiDocGetWithId, IndexName(userId), docId)
	content, err := SendRequest(ctx, api, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	md := &model.MeiliDocument{}
	err = json.Unmarshal(content, md)
	if err == nil && md.Code != "" {
		return nil, errors.New(md.Code)
	}

	return md, err
}

func PutDocument(ctx context.Context, userId int64, doc model.MeiliDocument) error {
	api := fmt.Sprintf(ApiDocCreate, IndexName(userId))
	content, err := SendRequest(ctx, api, http.MethodPut, doc)
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

func DelDocument(ctx context.Context, userId int64, docId string) error {
	api := fmt.Sprintf(ApiDocDelete, IndexName(userId), docId)
	content, err := SendRequest(ctx, api, http.MethodDelete, nil)
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
