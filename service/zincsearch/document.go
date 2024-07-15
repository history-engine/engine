package zincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"history-engine/engine/library/logger"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"io"
	"net/http"
)

// PutDocument 添加数据到ZincSearch索引
// TODO 请求ZincSearch逻辑封装
func PutDocument(userId int64, uniqueId string, doc *model.ZincDocument) error {
	body, _ := json.Marshal(doc)

	api := fmt.Sprintf(ApiDocCreateWithId, fmt.Sprintf("%s_%d", setting.ZincSearch.IndexPrefix, userId), uniqueId)
	req, _ := http.NewRequest("PUT", setting.ZincSearch.Host+api, bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "123456")

	res, err := client.Do(req)
	if err != nil || res == nil {
		logger.Zap().Error("put zinc index err", zap.Error(err))
		return err
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	logger.Zap().Debug("put zinc index", zap.String("status", res.Status), zap.String("body", string(body)))
	return nil
}
