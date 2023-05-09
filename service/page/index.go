package page

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		fmt.Printf("put zinc index err: %v\n", err)
		return err
	}
	defer res.Body.Close()

	body, _ = io.ReadAll(res.Body)
	fmt.Printf("put zinc index: %s\n", res.Status)
	return nil
}
