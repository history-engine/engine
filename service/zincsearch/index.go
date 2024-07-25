package zincsearch

import (
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
	_, err := SendRequest(ApiIndexCreateUpdateList, http.MethodPut, index)
	return err
}

func IndexName(userId int64) string {
	return fmt.Sprintf("%s_%s_%d", setting.ZincSearch.IndexPrefix, setting.Common.Env, userId)
}
