package page

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

// AddIndex 添加到zincsearch索引
func AddIndex(unique_id string, doc *model.ZincDocument) error {
	body, _ := json.Marshal(doc)

	api := fmt.Sprintf("%s/api/%s/_doc/%s", setting.ZincSearch.Host, setting.ZincSearch.Index, unique_id)
	req, _ := http.NewRequest("PUT", api, bytes.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("admin", "123456")

	res, err := http.DefaultClient.Do(req)
	if err != nil || res == nil {
		logger.Zap().Error("put zinc index err", zap.Error(err))
		return err
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	logger.Zap().Debug("put zinc index", zap.String("status", res.Status), zap.String("body", string(body)))
	return nil
}
