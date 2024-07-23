package zincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"io"
	"net/http"
)

// TODO 请求ZincSearch逻辑封装

// PutDocument 添加数据到ZincSearch索引
func PutDocument(userId int64, docId string, doc *model.ZincDocument) error {
	body, _ := json.Marshal(doc)
	api := fmt.Sprintf(ApiDocCreateWithId, fmt.Sprintf("%s_%d", setting.ZincSearch.IndexPrefix, userId), docId)
	req, _ := http.NewRequest("PUT", setting.ZincSearch.Host+api, bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil || res == nil {
		return err
	}

	_, err = io.ReadAll(res.Body)
	defer res.Body.Close()

	return err
}

// DelDocument 删除索引中的数据
func DelDocument(userId int64, uniqueId string, version int) error {
	index := fmt.Sprintf("%s_%d", setting.ZincSearch.IndexPrefix, userId)
	docId := fmt.Sprintf("%s%d", uniqueId, version)
	api := fmt.Sprintf(ApiDocDeleteWithId, index, docId)
	req, err := http.NewRequest(http.MethodDelete, setting.ZincSearch.Host+api, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(setting.ZincSearch.User, setting.ZincSearch.Password)

	res, err := client.Do(req)
	if err != nil || res == nil {
		return err
	}

	_, err = io.ReadAll(res.Body)
	defer res.Body.Close()

	return err
}
