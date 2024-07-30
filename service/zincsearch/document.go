package zincsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"net/http"
)

func GetDocument(ctx context.Context, userId int64, docId string) (doc model.ZincReadDocument, err error) {
	api := fmt.Sprintf(ApiDocGetWithId, IndexName(userId), docId)
	content, err := SendRequest(ctx, api, http.MethodGet, nil)
	if err != nil {
		return doc, err
	}

	zrd := model.ZincReadDocument{}
	err = json.Unmarshal(content, &zrd)
	if zrd.Error != "" {
		return doc, errors.New(zrd.Error)
	}

	return zrd, err
}

// PutDocument 添加数据到ZincSearch索引
func PutDocument(ctx context.Context, userId int64, docId string, doc *model.ZincWriteDocument) error {
	api := fmt.Sprintf(ApiDocCreateWithId, IndexName(userId), docId)
	content, err := SendRequest(ctx, api, http.MethodPut, doc)
	if err != nil {
		return err
	}

	zme := &model.ZincErrResp{}
	err = json.Unmarshal(content, zme)
	if zme.Error != "" {
		return errors.New(zme.Error)
	}

	return err
}

// DelDocument 删除索引中的数据
func DelDocument(ctx context.Context, userId int64, docId string) error {
	api := fmt.Sprintf(ApiDocDeleteWithId, IndexName(userId), docId)
	content, err := SendRequest(ctx, api, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	zme := &model.ZincErrResp{}
	err = json.Unmarshal(content, zme)
	if zme.Error != "" {
		return errors.New(zme.Error)
	}

	return err
}
