package zincsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"history-engine/engine/model"
	"history-engine/engine/setting"
	"net/http"
)

// CreateIndex 创建用户索引
func CreateIndex(userId int64) error {
	index := model.Index{
		Name:        IndexName(userId),
		StorageType: "disk",
		Settings:    &model.IndexSettings{},
		Mappings:    &model.Mappings{},
	}
	content, err := SendRequest(ApiIndexCreateUpdateList, http.MethodPut, index)
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

func IndexName(userId int64) string {
	return fmt.Sprintf("%s_%s_%d", setting.ZincSearch.IndexPrefix, setting.Common.Env, userId)
}
