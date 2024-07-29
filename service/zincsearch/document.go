package zincsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"net/http"
)

func GetDocument(userId int64, uniqueId string, version int) (*model.ZincReadDocument, error) {
	docId := fmt.Sprintf("%s%d", uniqueId, version)
	api := fmt.Sprintf(ApiDocGetWithId, IndexName(userId), docId)
	content, err := SendRequest(api, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	zrd := &model.ZincReadDocument{}
	err = json.Unmarshal(content, zrd)
	if zrd.Error != "" {
		return nil, errors.New(zrd.Error)
	}

	return zrd, err
}

// PutDocument 添加数据到ZincSearch索引
func PutDocument(userId int64, docId string, doc *model.ZincWriteDocument) error {
	api := fmt.Sprintf(ApiDocCreateWithId, IndexName(userId), docId)
	content, err := SendRequest(api, http.MethodPut, doc)
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
func DelDocument(userId int64, uniqueId string, version int) error {
	docId := fmt.Sprintf("%s%d", uniqueId, version)
	api := fmt.Sprintf(ApiDocDeleteWithId, IndexName(userId), docId)
	content, err := SendRequest(api, http.MethodDelete, nil)
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
