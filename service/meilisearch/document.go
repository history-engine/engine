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
	code, content, err := SendRequest(ctx, api, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	me := &model.MeiliError{}
	md := &model.MeiliDocument{}
	if code != 200 {
		err = json.Unmarshal(content, me)
		return nil, errors.New(me.Message)
	} else {
		err = json.Unmarshal(content, md)
		if err != nil {
			return nil, err
		}
	}

	return md, nil
}

func PutDocument(ctx context.Context, userId int64, doc model.MeiliDocument) error {
	api := fmt.Sprintf(ApiDocCreate, IndexName(userId))
	_, content, err := SendRequest(ctx, api, http.MethodPut, doc)
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
	_, content, err := SendRequest(ctx, api, http.MethodDelete, nil)
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
