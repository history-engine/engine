package zincsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"net/http"
)

// PutDocument 添加数据到ZincSearch索引
func PutDocument(userId int64, docId string, doc *model.ZincDocument) error {
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
